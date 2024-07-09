package main

import "strings"

func kebabToName(s string) string {
	words := strings.Split(s, "-")
	for i, word := range words {
		words[i] = strings.ToUpper(word[:1]) + word[1:]
	}
	return strings.Join(words, " ")
}
