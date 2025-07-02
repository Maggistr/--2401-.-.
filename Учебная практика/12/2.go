package main

import "fmt"

func findLongestString(strings ...string) string {
	if len(strings) == 0 {
		return ""
	}

	longest := strings[0]
	for _, s := range strings[1:] {
		if len(s) > len(longest) {
			longest = s
		}
	}
	return longest
}

func main() {
	longest := findLongestString(
		"Привет",
		"Hello, World!",
		"Go",
		"Это тестовая строка",
		"Еще одна строка",
	)

	fmt.Println("Самая длинная строка:", longest)

	fmt.Println(findLongestString("один", "два", "три", "четыре"))
}
