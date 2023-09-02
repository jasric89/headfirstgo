package main

import "fmt"

type CoffeeMachine struct {
	device Appliance
	pot    CoffeePot
}

type Appliance interface {
	TurnOn()
}

type MakeCoffee interface {
	Brew()
}

type Fan string

func (f Fan) TurnOn() {
	fmt.Println("Spinning")
}

type CoffeePot string

func (c CoffeePot) TurnOn() {
	fmt.Println("Powering Up")
}

func (c *CoffeePot) Brew() {
	fmt.Println("Heating Up")
}

func main() {
	var CoffeCup CoffeeMachine
	CoffeCup.device = Fan("Windco Breeze")
	CoffeCup.device.TurnOn()
	CoffeCup.device = CoffeePot("LuxBrew")
	CoffeCup.device.TurnOn()
	CoffeCup.pot = CoffeCup.device.(CoffeePot)
	CoffeCup.pot.Brew()
}
