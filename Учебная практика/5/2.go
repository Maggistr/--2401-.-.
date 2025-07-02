package main

import (
	"fmt"
	"strings"
)

func main() {
	sentence := "Go is an open source programming language that makes it easy to build simple, reliable, and efficient software."

	words := strings.Fields(sentence)

	fmt.Println("Слова в предложении:")
	for i, word := range words {
		fmt.Printf("%d: %s\n", i+1, word)
	}

	commaSentence := "apple,orange,banana,grape"
	fmt.Println("\nРазбиение строки по запятым:")
	fruits := strings.Split(commaSentence, ",")
	for _, fruit := range fruits {
		fmt.Println(fruit)
	}
}
