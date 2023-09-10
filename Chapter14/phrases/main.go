package main

import (
	"fmt"

	prose "github.com/jasric89/headfirstgo/chapter14/phrases"
)

func main() {
	phrases := []string{"myparents", "a rodeo clown"}
	fmt.Println("A Photo Of", prose.JoinWithCommas(phrases))
	phrases = []string{"my parents", "a rodeo cloud", "a prize ball"}
	fmt.Println("A Photo Of", prose.JoinWithCommas(phrases))
}
