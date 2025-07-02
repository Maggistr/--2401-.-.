package main

import (
	"fmt"
	"time"
)

type Student struct {
	Name      string
	BirthYear int
	Grades    []int
}

func (s Student) CalculateAge() int {
	currentYear := time.Now().Year()
	return currentYear - s.BirthYear
}

func (s Student) GetStatus() string {
	if len(s.Grades) == 0 {
		return "нет оценок"
	}

	sum := 0
	for _, grade := range s.Grades {
		sum += grade
	}
	average := float64(sum) / float64(len(s.Grades))

	switch {
	case average >= 4.5:
		return "отличник"
	case average >= 3.5:
		return "хорошист"
	default:
		return "троечник"
	}
}

func main() {
	student := Student{
		Name:      "Иван Иванов",
		BirthYear: 2000,
		Grades:    []int{5, 5, 4, 5, 4},
	}

	fmt.Printf("%s: возраст %d лет, статус - %s\n",
		student.Name,
		student.CalculateAge(),
		student.GetStatus())
}
