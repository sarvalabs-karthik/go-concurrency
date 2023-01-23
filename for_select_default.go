package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan interface{})
	go func() {
		time.Sleep(5 * time.Second)
		close(done)
	}()

	workCounter := 0
loop:
	for {
		select {
		case <-done:
			break loop
		default: // use default to avoid blocking and do other work in mean time
		}

		// Simulate work
		workCounter++
		time.Sleep(1 * time.Second)
	}

	fmt.Printf("achieved %v cycles of work before signalled to stop.\n", workCounter)
}

// go run for_select_default.go
