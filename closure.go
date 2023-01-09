package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	salutation := "hello"
	wg.Add(1)
	go func() {
		defer wg.Done()
		salutation = "world"
	}()
	wg.Wait()
	fmt.Println(salutation)
	printListWrongWay()
	printListRightWay()
}

func printListWrongWay() {
	var wg sync.WaitGroup
	for _, salutation := range []string{"apple", "bat", "cat"} {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(salutation)
		}()
	}
	wg.Wait()
	fmt.Println()
}

func printListRightWay() {
	var wg sync.WaitGroup
	for _, salutation := range []string{"apple", "bat", "cat"} {
		wg.Add(1)
		go func(salutation string) {
			defer wg.Done()
			fmt.Println(salutation)
		}(salutation)
	}

	wg.Wait()
}

// go run closure.go
