package main

import (
	"sync"
	"testing"
)

func BenchmarkContextSwitch(b *testing.B) {
	var wg sync.WaitGroup
	begin := make(chan struct{})
	c := make(chan struct{})

	var token struct{}
	sender := func() {
		defer wg.Done()
		<-begin // ignore cost of setting up and starting each goroutine.
		for i := 0; i < b.N; i++ {
			c <- token // equivalent to sending a message as it takes up no memory
		}
	}
	receiver := func() {
		defer wg.Done()
		<-begin // ignore cost of setting up and starting each goroutine.
		for i := 0; i < b.N; i++ {
			<-c // ignore received message
		}
	}

	wg.Add(2)
	go sender()
	go receiver()
	b.StartTimer() // start performance timing
	close(begin)   // start the two goroutines
	wg.Wait()
}
