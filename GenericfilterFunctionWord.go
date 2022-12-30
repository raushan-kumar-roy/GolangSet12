package main

import (
	"fmt"
)

func filter(slice []interface{}, f func(interface{}) bool) []interface{} {
	filtered := make([]interface{}, 0)
	for _, v := range slice {
		if f(v) {
			filtered = append(filtered, v)
		}
	}
	return filtered
}

func main() {
	words := []string{"ant", "beetle", "bee", "wasp", "butterfly", "fly", "moth", "beer"}
	wordsInterface := make([]interface{}, len(words))
	for i, v := range words {
		wordsInterface[i] = v
	}

	startsWithB := filter(wordsInterface, func(v interface{}) bool {
		return v.(string)[0] == 'b'
	})
	fmt.Println("Words starting with 'b':")
	fmt.Println(startsWithB)

	threeLetters := filter(wordsInterface, func(v interface{}) bool {
		return len(v.(string)) == 3
	})
	fmt.Println("3 character words:")
	fmt.Println(threeLetters)
}
