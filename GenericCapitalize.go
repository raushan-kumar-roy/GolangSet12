package main

import (
	"fmt"
	"strings"
)

type StringType string

func main() {
	words := []StringType{"ant", "beetle", "bee", "wasp", "butterfly", "fly", "moth"}

	capitalizedWords := Map(words, func(word StringType) StringType {
		return StringType(strings.Title(string(word)))
	})
	colonWords := Map(words, func(word StringType) StringType {
		return word + ":"
	})
	fmt.Println(capitalizedWords)
	fmt.Println(colonWords)
}

func Map(slice []StringType, fn func(StringType) StringType) []StringType {
	ModifiedStrings := make([]StringType, len(slice))
	for i, v := range slice {
		ModifiedStrings[i] = fn(v)
	}
	return ModifiedStrings
}
