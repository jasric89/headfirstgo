package main

import (
	"fmt"
)

func main() {

	var price int = 100
	fmt.Println("Price is", price, "dollars.")

	var taxRate float64 = 0.08
	var tax float64 = float64(price) * taxRate
	fmt.Println("Tax is", tax, "dollars.")

	var total float64 = float64(price) + tax
	fmt.Println("Total Cost is", total, "dollars.")

	var availableFunds int = 120
	fmt.Println(availableFunds, "Dollars available")
	fmt.Println("Within Budget?", int(total) <= availableFunds)

}
