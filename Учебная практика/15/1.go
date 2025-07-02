package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func main() {
	rect := Rectangle{Width: 5, Height: 3}
	circ := Circle{Radius: 2.5}

	shapes := []Shape{rect, circ}

	for _, shape := range shapes {
		fmt.Printf("Площадь фигуры: %.2f\n", shape.Area())
	}
}
