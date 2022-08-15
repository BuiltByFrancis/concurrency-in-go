package main

import "fmt"

/*
	OrDone is useful when you dont know if the channel you are reading from have been cancelled in another goroutine.
*/

func main() {
	orDone := func(done, c <-chan interface{}) <-chan interface{} {
		valStream := make(chan interface{})
		go func() {
			defer close(valStream)
			for {
				select {
				case <-done:
					return
				case v, ok := <-c:
					if !ok {
						return
					}
					select {
					case valStream <- v:
					case <-done:
					}
				}
			}
		}()
		return valStream
	}

	done := make(chan interface{})
	c := make(chan interface{})

	for val := range orDone(done, c) {
		fmt.Println(val)
	}
}
