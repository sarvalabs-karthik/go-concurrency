package main

import "fmt"

func main() {
	stringStream := make(chan string, 2)
	go func() {
		stringStream <- "hello world"
	}()
	fmt.Println(<-stringStream)
}

// go run channel.go
