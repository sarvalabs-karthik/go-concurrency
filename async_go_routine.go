package main

import (
	"fmt"
	"time"
)

// not synchronized
func main() {
	go func() {
		fmt.Println("some thing")
	}()
	go func() {
		fmt.Println("hello world")
	}()
	time.Sleep(10)
	go greet()
	//time.Sleep(1000)
}

func greet() {
	fmt.Println("welcome to mars")
}

// go run async_go_routine.go
