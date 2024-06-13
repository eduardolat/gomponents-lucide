package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

// minifySVGFolder minifies all the SVG's inside a directory using the minify binary
//
// See: https://github.com/tdewolff/minify/tree/master/cmd/minify
func minifySVGFolder(dirPath string) error {
	err := filepath.Walk(dirPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() {
			return nil
		}

		if filepath.Ext(path) != ".svg" {
			return nil
		}

		return minifySVG(path)
	})

	if err != nil {
		return fmt.Errorf("error walking the path %s: %w", dirPath, err)
	}

	return nil
}

// minifySVG minifies a SVG file using the minify binary
//
// See: https://github.com/tdewolff/minify/tree/master/cmd/minify
func minifySVG(filePath string) error {
	cmd := exec.Command("minify", "--type=svg", "-o", filePath, filePath)

	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("error minifying file %s: %w", filePath, err)
	}

	return nil
}
