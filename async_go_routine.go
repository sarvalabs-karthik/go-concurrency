package main

import (
	"fmt"
)

// not synchronized
func main() {
	go func() {
		fmt.Println("hello world")
	}()

	go greet()
	//time.Sleep(1000)
}

func greet() {
	fmt.Println("welcome to mars")
}

// go run async_go_routine.go
