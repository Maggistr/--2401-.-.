package main

import "fmt"

func main() {
	var dayNumber int

	fmt.Print("Введите номер дня недели (1-7): ")
	fmt.Scan(&dayNumber)

	var dayName string

	switch dayNumber {
	case 1:
		dayName = "Понедельник"
	case 2:
		dayName = "Вторник"
	case 3:
		dayName = "Среда"
	case 4:
		dayName = "Четверг"
	case 5:
		dayName = "Пятница"
	case 6:
		dayName = "Суббота"
	case 7:
		dayName = "Воскресенье"
	default:
		dayName = "Неверный номер дня недели"
	}

	fmt.Println(dayName)
}
