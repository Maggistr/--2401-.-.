package main

import "fmt"

func main() {
	grades := make(map[string]int)

	addGrade(grades, "Alice", 95)
	addGrade(grades, "Bob", 87)
	addGrade(grades, "Charlie", 91)

	fmt.Println("Alice's grade:", findGrade(grades, "Alice"))
	fmt.Println("Bob's grade:", findGrade(grades, "Bob"))

	fmt.Println("Dave's grade:", findGrade(grades, "Dave"))

	removeGrade(grades, "Bob")
	fmt.Println("After removal, Bob's grade:", findGrade(grades, "Bob"))

	fmt.Println("All grades:", grades)
}

func addGrade(grades map[string]int, name string, grade int) {
	grades[name] = grade
	fmt.Printf("Added grade for %s: %d\n", name, grade)
}

func findGrade(grades map[string]int, name string) int {
	grade, exists := grades[name]
	if !exists {
		return -1
	}
	return grade
}

func removeGrade(grades map[string]int, name string) {
	delete(grades, name)
	fmt.Printf("Removed grade for %s\n", name)
}
