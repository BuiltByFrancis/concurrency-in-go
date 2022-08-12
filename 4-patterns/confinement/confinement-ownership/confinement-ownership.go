package main

import (
	"fmt"
)

func main() {
	chanOwner := func() <-chan int {
		results := make(chan int, 5) // limit the scope of the write access for this channel to this anon method.
		go func() {
			defer close(results)
			for i := 0; i <= 5; i++ {
				results <- i
			}
		}()
		return results
	}

	consumer := func(results <-chan int) { // read-only copy of the channel
		for result := range results {
			fmt.Printf("Received: %d\n", result)
		}
		fmt.Println("Done receiving!")
	}

	results := chanOwner() // receive a read-only channel
	consumer(results)
}
