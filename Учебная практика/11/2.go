package main

import "fmt"

type Engine struct {
	Type         string
	Power        int     // в лошадиных силах
	Displacement float64 // объем в литрах
}

type Car struct {
	Make   string
	Model  string
	Year   int
	Engine Engine // вложенная структура
}

func NewCar(make, model string, year int, engineType string, power int, displacement float64) Car {
	return Car{
		Make:  make,
		Model: model,
		Year:  year,
		Engine: Engine{
			Type:         engineType,
			Power:        power,
			Displacement: displacement,
		},
	}
}

func (c Car) PrintInfo() {
	fmt.Printf("Автомобиль: %s %s (%d год)\n", c.Make, c.Model, c.Year)
	fmt.Printf("Двигатель: %s, %.1f л, %d л.с.\n",
		c.Engine.Type, c.Engine.Displacement, c.Engine.Power)
}

func (c Car) StartEngine() {
	fmt.Printf("Заводим %s %s... Двигатель %s работает!\n",
		c.Make, c.Model, c.Engine.Type)
}

func main() {
	car := NewCar("Toyota", "Camry", 2022, "V6", 301, 3.5)

	car.PrintInfo()
	fmt.Println()

	car.StartEngine()
	fmt.Println()

	sportsCar := NewCar("Porsche", "911", 2023, "Flat-6", 379, 3.0)
	sportsCar.PrintInfo()
	sportsCar.StartEngine()
}
