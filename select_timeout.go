package main

import (
	"fmt"
	"time"
)

func main() {
	var c <-chan int
	select {
	case <-c:
	case <-time.After(1 * time.Second):
		fmt.Println("timed out")
	}
}

// go run select_timeout.go
