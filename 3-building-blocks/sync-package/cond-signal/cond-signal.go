package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	c := sync.NewCond(&sync.Mutex{})    // <1> create condition
	queue := make([]interface{}, 0, 10) // <2> create queue of size 10

	removeFromQueue := func(delay time.Duration) {
		time.Sleep(delay)
		c.L.Lock()        // <8> enter critical section for the condition
		queue = queue[1:] // <9> simulate dequeing an item
		fmt.Println("Removed from queue")
		c.L.Unlock() // <10> exit critical section for condition
		c.Signal()   // <11> let the waiter know something has happened
	}

	for i := 0; i < 10; i++ {
		c.L.Lock()            // <3> enter critical section for the condition
		for len(queue) == 2 { // <4> ensure what we are waiting for has occured
			c.Wait() // <5> suspends goroutine until the signal on the condition has been set
		}
		fmt.Println("Adding to queue")
		queue = append(queue, struct{}{})
		go removeFromQueue(1 * time.Second) // <6> remove an element after one second
		c.L.Unlock()                        // <7> exit the conditions critical section
	}
}
