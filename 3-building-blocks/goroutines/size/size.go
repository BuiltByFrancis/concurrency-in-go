package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	memConsumed := func() uint64 {
		runtime.GC()
		var s runtime.MemStats
		runtime.ReadMemStats(&s)
		return s.Sys
	}

	var c <-chan interface{}
	var wg sync.WaitGroup
	noop := func() { wg.Done(); <-c } // This goroutine never exits, so we can see how much space each routine takes up.

	const numGoroutines = 1e4 // Large amount of routines to get a good average size of a goroutine.
	wg.Add(numGoroutines)
	before := memConsumed() // memory consumed before creating any routines
	for i := numGoroutines; i > 0; i-- {
		go noop()
	}
	wg.Wait()
	after := memConsumed() // memory consumed after creating all routines
	fmt.Printf("%.3fkb", float64(after-before)/numGoroutines/1000)
}
