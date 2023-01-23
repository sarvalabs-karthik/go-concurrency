package main

import (
	"fmt"
	"sync"
)

func main() {
	intStream := make(chan int)
	close(intStream)
	integer, ok := <-intStream
	fmt.Println(ok, integer)

	unBlockMultipleGoRoutines()
}

func unBlockMultipleGoRoutines() {
	begin := make(chan interface{})
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			<-begin
			fmt.Println("began ", i)
		}(i)
	}

	fmt.Println("unblocking go routines")
	close(begin)
	wg.Wait()
}

// go run channel_close.go
