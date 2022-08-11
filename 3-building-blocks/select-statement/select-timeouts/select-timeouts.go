package main

import (
	"fmt"
	"time"
)

func main() {
	var c <-chan int
	select {
	case <-c: // This case statement will never become unblocked
	case <-time.After(1 * time.Second):
		fmt.Println("Timed out.")
	}
}
