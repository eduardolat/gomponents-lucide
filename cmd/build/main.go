package main

import (
	"log"
	"os"
	"path"
	"strings"
)

const (
	version               = "0.452.0"
	iconsURL              = "https://github.com/lucide-icons/lucide/releases/download/" + version + "/lucide-icons-" + version + ".zip"
	repoIconsDir          = "https://raw.githubusercontent.com/lucide-icons/lucide/" + version + "/icons"
	tempDir               = "./tmp"
	iconsOutputFilePath   = "./lucide.go"
	infoOutputFilePath    = "./info.go"
	includedIconsFilePath = "./included_icons.txt"
)

func main() {
	os.RemoveAll(tempDir)
	err := os.MkdirAll(tempDir, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	defer os.RemoveAll(tempDir)

	// Download icons
	iconsFile := path.Join(tempDir, "icons.zip")
	if err := downloadFile(iconsFile, iconsURL); err != nil {
		log.Fatal(err)
	}

	// Unzip icons
	extractedDir := path.Join(tempDir, "extracted")
	if err := unzip(iconsFile, extractedDir); err != nil {
		log.Fatal(err)
	}

	// Minify SVG icons
	if err := minifySVGFolder(extractedDir); err != nil {
		log.Fatal(err)
	}

	// Read icons folder
	iconsDir := path.Join(extractedDir, "icons")
	files, err := os.ReadDir(iconsDir)
	if err != nil {
		log.Fatal(err)
	}

	// Generate Go code from icons
	components := []string{}
	infoVars := []string{}
	infoMaps := []string{}
	infoSlices := []string{}
	includedIcons := []string{}
	for _, file := range files {
		if file.IsDir() {
			continue
		}

		ext := path.Ext(file.Name())
		if ext != ".svg" && ext != ".json" {
			continue
		}

		kebabCaseName := strings.TrimSuffix(file.Name(), ext)
		funcName := kebabToFuncName(kebabCaseName)
		name := kebabToName(kebabCaseName)
		infoName := kebabToInfoName(kebabCaseName)

		filePath := path.Join(iconsDir, file.Name())
		b, err := os.ReadFile(filePath)
		if err != nil {
			log.Fatal(err)
		}

		if ext == ".svg" {
			component := generateComponent(file.Name(), funcName, b)
			components = append(components, component)
			includedIcons = append(includedIcons, funcName)
		}

		if ext == ".json" {
			infoVars = append(infoVars, generateInfoVars(file.Name(), name, kebabCaseName, funcName, infoName, b))
			infoMaps = append(infoMaps, generateInfoMap(kebabCaseName, infoName, name))
			infoSlices = append(infoSlices, generateInfoSlice(infoName))
		}
	}
	iconsFileContents := generateIconsFile(components)
	infoFileContents := generateInfoFile(infoVars, infoMaps, infoSlices)

	// Write icons Go code to file
	err = os.WriteFile(iconsOutputFilePath, iconsFileContents, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	// Write info Go code to file
	err = os.WriteFile(infoOutputFilePath, infoFileContents, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	// Write icons list to file
	includedIconsFileContents := strings.Join(includedIcons, "\n")
	err = os.WriteFile(includedIconsFilePath, []byte(includedIconsFileContents), os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("✅ Lucide icons generated successfully!")
}
