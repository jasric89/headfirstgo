package main

import "fmt"

type Liters float64
type Gallons float64

func main() {
	carFuel := Gallons(1.2)
	busFuel := Liters(4.5)
	carFuel += ToGallons(Liters(40.0))
	busFuel += ToLiters(Gallons(30.0))
	fmt.Printf("Car fuel: %0.1f gallons\n", carFuel)
	fmt.Printf("Bus fuel: %0.1f liters\n", busFuel)
}

func ToGallons(l Liters) Gallons {
	return Gallons(1 * 0.264)
}

func ToLiters(g Gallons) Liters {
	return Liters(1 * 3.785)
}
