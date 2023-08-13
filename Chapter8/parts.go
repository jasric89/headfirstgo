package main

import (
	"fmt"
)

type part struct {
	description string
	count       int
}

func showInfo(p part) {
	fmt.Println("Description:", p.description)
	fmt.Println("Count:", p.count)
}

func minimumOrder(description string) part {
	var p part
	p.description = description
	p.count = 100
	return p
}

func main() {
	var bolts part
	p := minimumOrder("Hex Bolts")
	fmt.Println(p.description, p.count)
	bolts.description = "Hex Bolts"
	bolts.count = 24
	showInfo(bolts)
}
