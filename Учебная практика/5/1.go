package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "Hello, World! This is Go Programming."

	length := len(str)
	fmt.Printf("1. Длина строки: %d символов\n", length)

	substr := "Go"
	contains := strings.Contains(str, substr)
	fmt.Printf("2. Строка содержит '%s': %t\n", substr, contains)

	upper := strings.ToUpper(str)
	lower := strings.ToLower(str)
	fmt.Println("3. В верхнем регистре:", upper)
	fmt.Println("   В нижнем регистре:", lower)

	fmt.Println("\nДополнительные операции:")
	fmt.Println("Количество слов 'is':", strings.Count(str, "is"))
	fmt.Println("Замена 'Go' на 'Golang':", strings.ReplaceAll(str, "Go", "Golang"))
	fmt.Println("Начинается с 'Hello':", strings.HasPrefix(str, "Hello"))
	fmt.Println("Заканчивается на '.':", strings.HasSuffix(str, "."))
}
