package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for _, salutation := range []string{"hello", "greetings", "good day"} {
		wg.Add(1)
		go func(salutation string) { // Pass in the variable
			defer wg.Done()
			fmt.Println(salutation)
		}(salutation) // no heap, no race conditions. (Order is still random as expected)
	}
	wg.Wait()
}
