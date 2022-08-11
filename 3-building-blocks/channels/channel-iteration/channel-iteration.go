package main

import (
	"fmt"
)

func main() {
	intStream := make(chan int)
	go func() {
		defer close(intStream) // Ensure the channel is closed before we exit the goroutine.
		for i := 1; i <= 5; i++ {
			intStream <- i
		}
	}()

	for integer := range intStream { // range over the channel
		fmt.Printf("%v ", integer)
	}
}
