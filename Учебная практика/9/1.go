package main

import "fmt"

func main() {
	var a, b float64
	var operator string

	fmt.Print("Введите первое число: ")
	fmt.Scan(&a)

	fmt.Print("Введите оператор (+, -, *, /): ")
	fmt.Scan(&operator)

	fmt.Print("Введите второе число: ")
	fmt.Scan(&b)

	var result float64
	var err error = nil

	switch operator {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		if b == 0 {
			err = fmt.Errorf("деление на ноль невозможно")
		} else {
			result = a / b
		}
	default:
		err = fmt.Errorf("неизвестный оператор")
	}

	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Printf("Результат: %.2f %s %.2f = %.2f\n", a, operator, b, result)
	}
}
