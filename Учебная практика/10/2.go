package main

import "fmt"

func isLeapYear(year int) bool {
	return (year%4 == 0 && year%100 != 0) || year%400 == 0
}

func main() {
	years := []int{2000, 2004, 2100, 2020, 2023}

	for _, year := range years {
		if isLeapYear(year) {
			fmt.Printf("%d - високосный год\n", year)
		} else {
			fmt.Printf("%d - не високосный год\n", year)
		}
	}
}
