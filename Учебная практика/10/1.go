package main

import "fmt"

func main() {
	a := 10
	b := 3

	fmt.Println("Арифметические операции:")
	fmt.Printf("%d + %d = %d\n", a, b, a+b)
	fmt.Printf("%d - %d = %d\n", a, b, a-b)
	fmt.Printf("%d * %d = %d\n", a, b, a*b)
	fmt.Printf("%d / %d = %d (целочисленное деление)\n", a, b, a/b)
	fmt.Printf("%d %% %d = %d (остаток от деления)\n", a, b, a%b)

	a++
	fmt.Println("После a++:", a)
	a--
	fmt.Println("После a--:", a)
}
