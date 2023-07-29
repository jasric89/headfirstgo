package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// guess challenges players to guess a random number between 1 and 100.
func main() {

	seconds := time.Now().Unix()
	rand.New(rand.NewSource(seconds))
	target := rand.Intn(100) + 1
	fmt.Println("Ive chosen a random number between 1 and 100.")
	fmt.Println("Can you guess It?")
	reader := bufio.NewReader(os.Stdin)
	success := false
	guesses := 0
	for guesses < 10 {
		fmt.Println("You Have", 10-guesses, "guesses left")
		fmt.Print("Make a Guess:")
		input, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)
		guess, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}
		if guess < target {
			fmt.Println("Oops. Your Guess Was LOW.")
		} else if guess > target {
			fmt.Println("Oops. Your Guess Was High!")
		} else {
			success = true
			fmt.Println("Good Job! You guessed it!")
			break
		}
		guesses++
	}
	if !success {
		fmt.Println("Sorry, you didnt guess my number. It was", target)
	}
}
