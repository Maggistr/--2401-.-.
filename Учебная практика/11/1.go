package main

import "fmt"

type Student struct {
	Name   string
	Age    int
	Course int
	GPA    float64
}

func NewStudent(name string, age, course int, gpa float64) Student {
	return Student{
		Name:   name,
		Age:    age,
		Course: course,
		GPA:    gpa,
	}
}

func (s Student) PrintInfo() {
	fmt.Printf("Студент: %s\nВозраст: %d\nКурс: %d\nСредний балл: %.2f\n",
		s.Name, s.Age, s.Course, s.GPA)
}

func (s *Student) Promote() {
	s.Course++
	fmt.Printf("%s переведен на %d курс\n", s.Name, s.Course)
}

func (s *Student) UpdateGPA(newGPA float64) {
	s.GPA = newGPA
	fmt.Printf("Средний балл %s обновлен: %.2f\n", s.Name, s.GPA)
}

func main() {
	student := NewStudent("Иван Иванов", 20, 2, 4.5)

	student.PrintInfo()
	fmt.Println()

	student.Promote()

	student.UpdateGPA(4.7)
	fmt.Println()

	student.PrintInfo()
}
