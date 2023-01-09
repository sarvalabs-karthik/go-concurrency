package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	sayHello := func() {
		defer wg.Done()
		fmt.Println("hello world")
	}
	wg.Add(1)
	go sayHello()
	wg.Wait()
}

// go run join.go
