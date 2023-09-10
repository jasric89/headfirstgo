package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func responseSize(url string, channel chan int) {
	fmt.Println("Getting url", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	channel <- len(body)
}

func main() {
	sizes := make(chan int)
	go responseSize("https://engadget.com", sizes)
	go responseSize("https://golang.org", sizes)
	go responseSize("https://golang.org/doc", sizes)
	fmt.Println(<-sizes)
	fmt.Println(<-sizes)
	fmt.Println(<-sizes)
}
