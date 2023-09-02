package main

import "fmt"

type Liters float64
type Gallons float64
type Millilitres float64

func main() {
	soda := Liters(2)
	fmt.Printf("%0.3f liters equals %0.3f gallons\n", soda, soda.ToGallons())
	water := Millilitres(500)
	fmt.Printf("%0.3f milliters equals %0.3f gallons\n", water, water.ToGallons())
}

func (l Liters) ToGallons() Gallons {
	return Gallons(l * 0.264)
}

func (m Millilitres) ToGallons() Gallons {
	return Gallons(m * 0.000264)
}
