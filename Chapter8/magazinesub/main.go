package main

import (
	"fmt"

	"github.com/jasric89/headfirstgo/Chapter8/magazineone"
)

func main() {
	var s magazineone.Subscriber
	s.rate = 4.99
	fmt.Println(s.rate)
}
