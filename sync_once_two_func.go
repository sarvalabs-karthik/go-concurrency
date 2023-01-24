package main

import (
	"fmt"
	"sync"
)

func main() {
	var count int

	increment := func() { count++ }
	decrement := func() { count-- }

	var once sync.Once

	once.Do(increment)
	once.Do(decrement)

	fmt.Println("count : ", count)
}

//  go run sync_once_two_func.go
