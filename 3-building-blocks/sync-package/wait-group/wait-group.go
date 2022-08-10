package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1) // indicates that 1 goroutine is beginning
	go func() {
		defer wg.Done() // called just before exiting the function to tell the wait group we are finished
		fmt.Println("1st goroutine sleeping...")
		time.Sleep(1 * time.Nanosecond)
	}()

	wg.Add(1) // indicates that a second goroutine is beginning
	go func() {
		defer wg.Done() // called just before exiting the function to tell the wait group we are finished
		fmt.Println("2nd goroutine sleeping...")
		time.Sleep(2 * time.Nanosecond)
	}()

	wg.Wait() // blocks the main goroutine until all goroutines have exited.
	fmt.Println("All goroutines complete.")
}
