package main

import "fmt"

type Transport interface {
	Move()
	Stop()
}

type Car struct {
	Model string
}

func (c Car) Move() {
	fmt.Printf("Автомобиль %s начал движение\n", c.Model)
}

func (c Car) Stop() {
	fmt.Printf("Автомобиль %s остановился\n", c.Model)
}

type Bicycle struct {
	Brand string
}

func (b Bicycle) Move() {
	fmt.Printf("Велосипед %s начал движение\n", b.Brand)
}

func (b Bicycle) Stop() {
	fmt.Printf("Велосипед %s остановился\n", b.Brand)
}

type Train struct {
	Number int
}

func (t Train) Move() {
	fmt.Printf("Поезд №%d начал движение\n", t.Number)
}

func (t Train) Stop() {
	fmt.Printf("Поезд №%d остановился\n", t.Number)
}

func main() {
	vehicles := []Transport{
		Car{Model: "Toyota Camry"},
		Bicycle{Brand: "Stels"},
		Train{Number: 123},
	}

	for _, vehicle := range vehicles {
		vehicle.Move()
		vehicle.Stop()
		fmt.Println("---")
	}
}
