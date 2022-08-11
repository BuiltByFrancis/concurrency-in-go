package main

import (
	"fmt"
)

func main() {
	intStream := make(chan int)
	close(intStream)
	integer, ok := <-intStream // ok returns false, integer returns 0 which is the default value
	fmt.Printf("(%v): %v", ok, integer)
}
