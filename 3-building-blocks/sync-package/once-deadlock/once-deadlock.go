package main

import (
	"sync"
)

func main() {
	var onceA, onceB sync.Once
	var initB func()
	initA := func() { onceB.Do(initB) }
	initB = func() { onceA.Do(initA) } // cant proceed until call below returns
	onceA.Do(initA)                    // will never return due to the circular reference
}
