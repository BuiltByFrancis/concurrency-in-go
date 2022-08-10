package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	salutation := "hello"
	wg.Add(1)
	go func() {
		defer wg.Done()
		salutation = "welcome" // goroutines execute within the same address space they were created in
	}()
	wg.Wait()
	fmt.Println(salutation) // Prints welcome, NOT hello.
}
