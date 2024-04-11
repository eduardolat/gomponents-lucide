package main

import "strings"

func kebabToUpperCamelCase(s string) string {
	var result string
	words := strings.Split(s, "-")
	for _, word := range words {
		if word == "" {
			continue
		}
		result += strings.ToUpper(word[:1]) + word[1:]
	}
	return result
}
