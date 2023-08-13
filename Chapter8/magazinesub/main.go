package main

import (
	"fmt"

	"github.com/jasric89/headfirstgo/Chapter8/magazineone"
)

func main() {
	var s magazineone.Subscriber
	s.Rate = 4.99
	fmt.Println(s.Rate)
}
