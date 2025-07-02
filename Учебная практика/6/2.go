package main

import "fmt"

func removeElement(slice []string, index int) []string {
	if index < 0 || index >= len(slice) {
		fmt.Println("Неверный индекс")
		return slice
	}

	return append(slice[:index], slice[index+1:]...)
}

func main() {

	mySlice := []string{"First", "Second", "Third", "Fourth"}

	fmt.Println("Исходный срез:", mySlice)

	mySlice = removeElement(mySlice, 2)

	fmt.Println("Срез после удаления:", mySlice)
}
