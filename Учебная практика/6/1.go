package main

import "fmt"

func main() {
	var mySlice []string

	mySlice = append(mySlice, "First")
	mySlice = append(mySlice, "Second")
	mySlice = append(mySlice, "Third")

	fmt.Println("Элементы среза:")
	for i, element := range mySlice {
		fmt.Printf("Индекс: %d, Значение: %s\n", i, element)
	}
}
