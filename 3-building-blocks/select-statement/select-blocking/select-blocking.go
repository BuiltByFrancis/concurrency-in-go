package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	c := make(chan interface{})
	go func() {
		time.Sleep(5 * time.Second)
		close(c) // wait 5 seconds and then close the channel
	}()

	fmt.Println("Blocking on read...")
	select {
	case <-c: // attempt to read the channel, as warning suggests this is not required, just demonstrates this example.
		fmt.Printf("Unblocked %v later.\n", time.Since(start))
	}
}
