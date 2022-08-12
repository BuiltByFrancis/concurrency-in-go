package main

import (
	"fmt"
	"time"
)

func main() {
	var or func(channels ...<-chan interface{}) <-chan interface{}

	/*
		This function allows you to combine any number of channels together into a single channel
		that will close when any one of its component channels is closed.
	*/
	or = func(channels ...<-chan interface{}) <-chan interface{} {
		switch len(channels) {
		case 0: // termination from recursion
			return nil
		case 1: // termination from recursion
			return channels[0]
		}

		orDone := make(chan interface{})
		go func() { // go routine to wait without blocking
			defer close(orDone)

			switch len(channels) {
			case 2: // all calls here have at least 2 channels
				select {
				case <-channels[0]:
				case <-channels[1]:
				}
			default:
				select {
				case <-channels[0]:
				case <-channels[1]:
				case <-channels[2]:
				case <-or(append(channels[3:], orDone)...): // recurse, this creates a tree like structure
				}
			}
		}()
		return orDone
	}

	sig := func(after time.Duration) <-chan interface{} { // creates a channel that will close after a given time frame
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	start := time.Now()
	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)
	fmt.Printf("done after %v", time.Since(start)) // Should see the or channel closes after the shortest time (1 second).
}
