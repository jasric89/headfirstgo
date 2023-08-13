package main

import "fmt"

type car struct {
	name     string
	topspeed float64
}

func nitroBoost(c *car) {
	c.topspeed += 50
}

func main() {
	var mustang car
	mustang.name = "Mustang Cobra"
	mustang.topspeed = 225
	nitroBoost(&mustang)
	fmt.Println(mustang.name)
	fmt.Println(mustang.topspeed)
}
