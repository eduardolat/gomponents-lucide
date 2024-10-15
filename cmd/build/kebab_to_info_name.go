package main

import "strings"

func kebabToInfoName(s string) string {
	var result string
	words := strings.Split(s, "-")
	for i, word := range words {
		if word == "" {
			continue
		}
		if i == 0 {
			result += word
			continue
		}
		result += strings.ToUpper(word[:1]) + word[1:]
	}
	return result + `Info`
}
