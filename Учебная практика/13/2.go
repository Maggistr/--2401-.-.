package main

import "fmt"

func incrementByValue(a int) {
	a = a + 1
	fmt.Println("Внутри incrementByValue:", a)
}

func incrementByPointer(a *int) {
	*a = *a + 1
	fmt.Println("Внутри incrementByPointer:", *a)
}

func main() {
	val := 10

	fmt.Println("Исходное значение:", val)

	incrementByValue(val)
	fmt.Println("После incrementByValue:", val)

	incrementByPointer(&val)
	fmt.Println("После incrementByPointer:", val)

	fmt.Println("\nРазница между передачей по значению и по указателю:")
	fmt.Println("Передача по значению создает копию, оригинал не меняется")
	fmt.Println("Передача по указателю позволяет изменять оригинальную переменную")
}
