package main

import (
	"fmt"
	"sort"
)

func findElement(slice []int, target int) (int, bool) {
	for i, v := range slice {
		if v == target {
			return i, true
		}
	}
	return -1, false
}

func sortSlice(slice []int) []int {
	sorted := make([]int, len(slice))
	copy(sorted, slice)
	sort.Ints(sorted)
	return sorted
}

func filterSlice(slice []int) []int {
	var filtered []int
	for _, v := range slice {
		if v%2 == 0 {
			filtered = append(filtered, v)
		}
	}
	return filtered
}

func main() {
	numbers := []int{5, 3, 8, 1, 2, 7, 4, 6}

	if index, found := findElement(numbers, 8); found {
		fmt.Printf("Элемент 8 найден по индексу %d\n", index)
	} else {
		fmt.Println("Элемент не найден")
	}

	sorted := sortSlice(numbers)
	fmt.Println("Отсортированный срез:", sorted)

	filtered := filterSlice(numbers)
	fmt.Println("Отфильтрованный срез (четные числа):", filtered)
}
