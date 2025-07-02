package main

import "fmt"

func main() {
	var i int = -42
	var i8 int8 = -128
	var i16 int16 = 32767
	var i32 int32 = -2147483648
	var i64 int64 = 9223372036854775807
	var ui uint = 42
	var ui8 uint8 = 255
	var ui16 uint16 = 65535
	var ui32 uint32 = 4294967295
	var ui64 uint64 = 18446744073709551615
	var r rune = 'Я'
	var b byte = 255

	var f32 float32 = 3.14
	var f64 float64 = 3.141592653589793

	var c64 complex64 = 1 + 2i
	var c128 complex128 = 3 + 4i

	var isTrue bool = true
	var isFalse bool = false

	var s string = "Hello, Go!"
	var multiLine string = `Многострочная
строка`

	fmt.Println("Целые числа:")
	fmt.Println("int:", i)
	fmt.Println("int8:", i8)
	fmt.Println("int16:", i16)
	fmt.Println("int32:", i32)
	fmt.Println("int64:", i64)
	fmt.Println("uint:", ui)
	fmt.Println("uint8:", ui8)
	fmt.Println("uint16:", ui16)
	fmt.Println("uint32:", ui32)
	fmt.Println("uint64:", ui64)
	fmt.Println("rune:", r, string(r))
	fmt.Println("byte:", b)

	fmt.Println("\nЧисла с плавающей точкой:")
	fmt.Println("float32:", f32)
	fmt.Println("float64:", f64)

	fmt.Println("\nКомплексные числа:")
	fmt.Println("complex64:", c64)
	fmt.Println("complex128:", c128)

	fmt.Println("\nЛогические значения:")
	fmt.Println("bool true:", isTrue)
	fmt.Println("bool false:", isFalse)

	fmt.Println("\nСтроки:")
	fmt.Println("string:", s)
	fmt.Println("многострочная string:", multiLine)
}
