package main

import (
	"fmt"
	"math"
)

func main() {
	const pi = math.Pi
	const e = math.E
	const goldenRatio = 1.61803398875

	radius := 5.0
	circleArea := pi * math.Pow(radius, 2)
	fmt.Printf("Area of circle with radius %.2f: %.2f\n", radius, circleArea)

	x := 2.0
	exp := math.Pow(e, x)
	fmt.Printf("e^%.2f = %.4f\n", x, exp)

	a := 10.0
	b := a / goldenRatio
	fmt.Printf("For length %.2f, golden ratio gives %.2f\n", a, b)

	const sqrt2 = math.Sqrt2
	const ln2 = math.Ln2

	fmt.Println("Square root of 2:", sqrt2)
	fmt.Println("Natural logarithm of 2:", ln2)
}
