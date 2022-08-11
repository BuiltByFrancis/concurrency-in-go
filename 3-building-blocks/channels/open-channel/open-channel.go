package main

import (
	"fmt"
)

func main() {
	stringStream := make(chan string)
	go func() {
		stringStream <- "Hello channels!"
	}()
	salutation, ok := <-stringStream // ok: true = written value, false = default value e.g. 0 for an int.
	fmt.Printf("(%v): %v", ok, salutation)
}
