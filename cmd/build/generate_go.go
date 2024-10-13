package main

import (
	"encoding/json"
	"fmt"
	"go/format"
	"log"
	"strings"
)

type IconInfo struct {
	Tags       []string `json:"tags"`
	Categories []string `json:"categories"`
}

func generateComponent(fileName string, funcName string, svgBytes []byte) string {
	fullSvg := string(svgBytes)

	// Find where the children starts
	start := strings.Index(fullSvg, ">") + 1
	if start == -1 {
		log.Fatalf("could not find the start of the svg tag in %s", fileName)
	}

	// Remove the svg tag and format the svg
	svg := fullSvg[start:]
	svg = strings.ReplaceAll(svg, "</svg>", "")
	svg = strings.TrimSpace(svg)

	previewURL := repoIconsDir + "/" + fileName

	fn := `
		// ` + funcName + ` icon: ` + previewURL + `
		func ` + funcName + `(children ...gomponents.Node) gomponents.Node {
			return svgWrapper(
				gomponents.Group(children),
				gomponents.Raw(` + "`" + svg + "`" + `),
			)
		}
	`

	return fn
}

func generateInfo(fileName string, name string, funcName string, jsonBytes []byte) string {
	var info IconInfo
	if err := json.Unmarshal(jsonBytes, &info); err != nil {
		log.Fatalf("could not unmarshal %s: %v", fileName, err)
	}

	tpl := `{
		Name: "%s",
		Icon: %s,
		Tags: []string{%s},
		Categories: []string{%s},
	},`

	tags := ""
	for _, tag := range info.Tags {
		tags += `"` + tag + `",`
	}

	categories := ""
	for _, category := range info.Categories {
		categories += `"` + category + `",`
	}

	return fmt.Sprintf(tpl, name, funcName, tags, categories)
}

func generatePackageDef() string {
	return `
		// Code generated by lucide-icons build task. DO NOT EDIT.
		// v` + version + `

		package lucide
	`
}

func generateIconsFile(components []string) []byte {
	pkg := generatePackageDef() + `

		import "maragu.dev/gomponents"

		// svgWrapper just creates the svg skeleton following the lucide
		// guidelines. It is used by all the icons.
		//
		// It includes an extra attribute data-glucide="icon" to globally
		// identify the icons generated by this package. It can be used to
		// style all the icons at once using CSS.
		//
		// Same as:
		//
		//	<svg
		//		xmlns="http://www.w3.org/2000/svg"
		//		width="24"
		//		height="24"
		//		viewBox="0 0 24 24"
		//		fill="none"
		//		stroke="currentColor"
		//		stroke-width="2"
		//		stroke-linecap="round"
		//		stroke-linejoin="round"
		//		data-glucide="icon"
		//	></svg>
		func svgWrapper(children ...gomponents.Node) gomponents.Node {
			return gomponents.El(
				"svg",
				gomponents.Attr("xmlns", "http://www.w3.org/2000/svg"),
				gomponents.Attr("width", "24"),
				gomponents.Attr("height", "24"),
				gomponents.Attr("viewBox", "0 0 24 24"),
				gomponents.Attr("fill", "none"),
				gomponents.Attr("stroke", "currentColor"),
				gomponents.Attr("stroke-width", "2"),
				gomponents.Attr("stroke-linecap", "round"),
				gomponents.Attr("stroke-linejoin", "round"),
				gomponents.Attr("data-glucide", "icon"),
				gomponents.Group(children),
			)
		}
	`
	pkg += strings.Join(components, "\n\n")

	b, err := format.Source([]byte(pkg))
	if err != nil {
		log.Fatal(err)
	}

	return b
}

func generateInfoFile(infos []string) []byte {
	pkg := generatePackageDef() + `

		import "maragu.dev/gomponents"

		// IconInfo represents the information of an icon.
		type IconInfo struct {
			Name       string
			Icon       func (children ...gomponents.Node) gomponents.Node
			Tags       []string
			Categories []string
		}

		// IconsInfo is a list of all the icons information.
		var IconsInfo = []IconInfo{
	`
	pkg += strings.Join(infos, "\n") + "\n"
	pkg += `}`

	b, err := format.Source([]byte(pkg))
	if err != nil {
		log.Fatal(err)
	}

	return b
}
