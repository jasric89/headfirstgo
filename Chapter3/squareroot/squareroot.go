package main

import (
	"fmt"
	"math"
)

func main() {
	root, err := squareRoot(9.3)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%0.3f", root)
	}
}

func squareRoot(number float64) (float64, error) {
	if number < 0 {
		return 0, fmt.Errorf("cant get square root of negative number")
	}
	return math.Sqrt(number), nil
}
