package main

import (
	"fmt"

	prose "phrases/prose"
)

func main() {
	phrases := []string{"myparents"}
	fmt.Println("A Photo Of", prose.JoinWithCommas(phrases))
	phrases = []string{"myparents", "a rodeo clown"}
	fmt.Println("A Photo Of", prose.JoinWithCommas(phrases))
	phrases = []string{"my parents", "a rodeo cloud", "a prize ball"}
	fmt.Println("A Photo Of", prose.JoinWithCommas(phrases))
}
