package main

import (
	"fmt"
	"sync"
	"time"
)

/*
	This example code will throw: "fatal error: all goroutines are asleep - deadlock!"

	To detect, prevent and correct deadlocks, use the coffman conditions:

	Mutual Exclusion:
		A concurrent process holds exclusive rights to a resource at any one time.
	Wait for condition:
		A concurrent process must simultaneously hold a resouirce and be waiting for an additional resource.
	No Preemption:
		A resource held by a concurrent process can only be released by that process, so it fulfills this condition.
	Circular Wait:
		A concurrent process (P1) must be waiting on a chain of other concurrent processes (P2), which are in turn
		waiting on it (P1), so it fulfills this final condition too.
*/

func main() {
	type value struct {
		mu    sync.Mutex
		value int
	}

	var wg sync.WaitGroup
	printSum := func(v1, v2 *value) {
		defer wg.Done()
		v1.mu.Lock()
		defer v1.mu.Unlock()

		// Ironically this is also a race condition.. lul
		time.Sleep(2 * time.Second)
		v2.mu.Lock()
		defer v2.mu.Unlock()

		fmt.Printf("sum=%v\n", v1.value+v2.value)
	}

	var a, b value
	wg.Add(2)
	go printSum(&a, &b)
	go printSum(&b, &a)
	wg.Wait()
}
