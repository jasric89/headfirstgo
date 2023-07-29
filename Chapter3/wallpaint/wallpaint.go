package main

import (
	"fmt"
	"log"
)

func main() {
	var width, height float64
	width = 4.2
	height = 3.0
	amount, err := paintNeeded(width, height)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%0.2f litres needed\n", amount)
}

func paintNeeded(width float64, height float64) (float64, error) {
	if width < 0 || height < 0 {
		return 0, fmt.Errorf("can not accept a number less than 0")
	}
	area := width * height
	return area / 10.0, nil
}
