package main

import (
	"fmt"
)

func main() {
	stringStream := make(chan string)
	go func() {
		stringStream <- "Hello channels!" // pass string literal into the channel.
	}()
	fmt.Println(<-stringStream) // read the string literal from the channel. This call is blocking
}
