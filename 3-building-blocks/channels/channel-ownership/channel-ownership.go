package main

import (
	"fmt"
)

func main() {
	chanOwner := func() <-chan int {
		resultStream := make(chan int, 5) // owner creates buffered channel
		go func() {                       // perform writes on the channel
			defer close(resultStream) // close the channel when the go routine exits
			for i := 0; i <= 5; i++ {
				resultStream <- i
			}
		}()
		return resultStream // return read only channel
	}

	resultStream := chanOwner()
	for result := range resultStream { // range over the channel
		fmt.Printf("Received: %d\n", result)
	}
	fmt.Println("Done receiving!")
}
