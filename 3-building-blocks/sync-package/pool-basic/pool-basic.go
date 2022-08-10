package main

import (
	"fmt"
	"sync"
)

func main() {
	myPool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("Creating new instance.")
			return struct{}{}
		},
	}

	myPool.Get()             // Creates a new instance since the pool is empty
	instance := myPool.Get() // creates a new instance since the pool is still empty
	myPool.Put(instance)     // returns the instance to the pool
	myPool.Get()             // does not need to create a new instance
}
