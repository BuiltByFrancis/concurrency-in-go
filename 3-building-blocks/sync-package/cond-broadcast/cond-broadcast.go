package main

import (
	"fmt"
	"sync"
)

func main() {
	type Button struct { // simulated button type
		Clicked *sync.Cond
	}
	button := Button{Clicked: sync.NewCond(&sync.Mutex{})}

	subscribe := func(c *sync.Cond, fn func()) { // A function that will allow us to register functions to handle signals from a condition.
		var goroutineRunning sync.WaitGroup // Each handler will run on its own goroutine, and subscribe will not exit until that goroutine is confirmed to be running
		goroutineRunning.Add(1)
		go func() {
			goroutineRunning.Done()
			c.L.Lock()
			defer c.L.Unlock()
			c.Wait()
			fn()
		}()
		goroutineRunning.Wait()
	}

	var clickRegistered sync.WaitGroup // ensures program doesnt exit before our prints occur
	clickRegistered.Add(3)
	subscribe(button.Clicked, func() {
		fmt.Println("Maximizing window.")
		clickRegistered.Done()
	})
	subscribe(button.Clicked, func() {
		fmt.Println("Displaying annoying dialogue box!")
		clickRegistered.Done()
	})
	subscribe(button.Clicked, func() {
		fmt.Println("Mouse clicked.")
		clickRegistered.Done()
	})

	button.Clicked.Broadcast() // tells all subscriptions something has happened, so the wait on 21 continutes to run the handler functions.

	clickRegistered.Wait()
}
