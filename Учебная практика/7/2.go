package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	text := `Go is an open source programming language that makes it easy to build simple, 
    reliable, and efficient software. Go is a statically typed, compiled programming 
    language designed at Google.`

	wordFrequency := countWords(text)

	fmt.Println("Word frequency:")
	for word, count := range wordFrequency {
		fmt.Printf("%s: %d\n", word, count)
	}
}

func countWords(text string) map[string]int {
	frequency := make(map[string]int)

	words := strings.FieldsFunc(text, func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	})

	for _, word := range words {
		lowerWord := strings.ToLower(word)
		frequency[lowerWord]++
	}

	return frequency
}
