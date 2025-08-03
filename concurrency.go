package main

import (
	"fmt"
	"time"
)

func main() {
	counter := 0

	// Two goroutines are writing to the same variable concurrently (race condition)
	go func() {
		for i := 0; i < 1000; i++ {
			counter++
		}
	}()

	go func() {
		for i := 0; i < 1000; i++ {
			counter++
		}
	}()

	// Short wait to allow goroutines to complete
	time.Sleep(1 * time.Millisecond)

	// Expected value is 2000, but due to the race condition it will usually be smaller
	fmt.Println("Counter:", counter)
}
