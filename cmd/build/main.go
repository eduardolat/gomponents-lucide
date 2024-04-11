package main

import (
	"log"
	"os"
	"path"
	"strings"
)

const version = "0.367.0"
const tempDir = "./tmp"
const outputFilePath = "./lucide.go"
const iconsURL = "https://github.com/lucide-icons/lucide/releases/download/0.367.0/lucide-icons-0.367.0.zip"

func main() {
	os.RemoveAll(tempDir)
	err := os.MkdirAll(tempDir, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	// defer os.RemoveAll(tempDir)

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

	// Read icons folder
	iconsDir := path.Join(extractedDir, "icons")
	files, err := os.ReadDir(iconsDir)
	if err != nil {
		log.Fatal(err)
	}

	// Generate Go code from icons
	components := []string{}
	for _, file := range files {
		if file.IsDir() {
			continue
		}

		ext := path.Ext(file.Name())
		if ext != ".svg" {
			continue
		}

		kebabCaseName := strings.TrimSuffix(file.Name(), ext)
		upperCamelCaseName := kebabToUpperCamelCase(kebabCaseName)

		filePath := path.Join(iconsDir, file.Name())
		b, err := os.ReadFile(filePath)
		if err != nil {
			log.Fatal(err)
		}

		component := generateComponent(file.Name(), upperCamelCaseName, b)
		components = append(components, component)
	}
	pkg := generatePackage(components)

	// Write Go code to file
	err = os.WriteFile(outputFilePath, []byte(pkg), os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("✅ Lucide icons generated successfully!")
}
