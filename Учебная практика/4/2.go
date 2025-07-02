package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var i int
	var i8 int8
	var i16 int16
	var i32 int32
	var i64 int64
	var ui uint
	var ui8 uint8
	var ui16 uint16
	var ui32 uint32
	var ui64 uint64
	var r rune
	var b byte

	var f32 float32
	var f64 float64

	var c64 complex64
	var c128 complex128

	var bl bool

	var s string

	fmt.Println("Размеры типов в байтах:")
	fmt.Println("int:", unsafe.Sizeof(i))
	fmt.Println("int8:", unsafe.Sizeof(i8))
	fmt.Println("int16:", unsafe.Sizeof(i16))
	fmt.Println("int32:", unsafe.Sizeof(i32))
	fmt.Println("int64:", unsafe.Sizeof(i64))
	fmt.Println("uint:", unsafe.Sizeof(ui))
	fmt.Println("uint8:", unsafe.Sizeof(ui8))
	fmt.Println("uint16:", unsafe.Sizeof(ui16))
	fmt.Println("uint32:", unsafe.Sizeof(ui32))
	fmt.Println("uint64:", unsafe.Sizeof(ui64))
	fmt.Println("rune:", unsafe.Sizeof(r))
	fmt.Println("byte:", unsafe.Sizeof(b))

	fmt.Println("\nЧисла с плавающей точкой:")
	fmt.Println("float32:", unsafe.Sizeof(f32))
	fmt.Println("float64:", unsafe.Sizeof(f64))

	fmt.Println("\nКомплексные числа:")
	fmt.Println("complex64:", unsafe.Sizeof(c64))
	fmt.Println("complex128:", unsafe.Sizeof(c128))

	fmt.Println("\nЛогический тип:")
	fmt.Println("bool:", unsafe.Sizeof(bl))

	fmt.Println("\nСтроки:")
	fmt.Println("string:", unsafe.Sizeof(s), "(размер структуры строки)")

	exampleString := "пример"
	fmt.Println("Длина строки 'пример' в байтах:", len(exampleString))
}
