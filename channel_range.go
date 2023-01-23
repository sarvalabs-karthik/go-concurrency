package main

import "fmt"

func main() {
	intStream := make(chan int)
	go func() {
		defer close(intStream)
		for i := 0; i < 5; i++ {
			intStream <- i
		}
	}()

	for i := range intStream {
		fmt.Println(i)
	}
}

// go run channel_range.go
