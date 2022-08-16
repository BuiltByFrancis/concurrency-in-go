package main

import (
	"fmt"
	"time"
)

func main() {
	doWork := func(
		done <-chan interface{},
		pulseInterval time.Duration,
	) (<-chan interface{}, <-chan time.Time) {
		heartbeat := make(chan interface{}) // hearbeat channel
		results := make(chan time.Time)
		go func() {
			defer close(heartbeat)
			defer close(results)

			pulse := time.Tick(pulseInterval)       // setup the pulse interval
			workGen := time.Tick(2 * pulseInterval) // simulate work after 2 pulses

			sendPulse := func() {
				select {
				case heartbeat <- struct{}{}:
				default: // guard against no one listening
				}
			}
			sendResult := func(r time.Time) {
				for {
					select {
					case <-done:
						return
					case <-pulse: // all send and recieve sections should include a pulse just like a done
						sendPulse()
					case results <- r:
						return
					}
				}
			}

			for {
				select {
				case <-done:
					return
				case <-pulse: // all send and recieve sections should include a pulse just like a done
					sendPulse()
				case r := <-workGen:
					sendResult(r)
				}
			}
		}()
		return heartbeat, results
	}
	done := make(chan interface{})
	time.AfterFunc(10*time.Second, func() { close(done) }) // Standard done channel to run this example for 10 seconds

	const timeout = 2 * time.Second               // set our timeout period, if we get nothing after this time, stop.
	heartbeat, results := doWork(done, timeout/2) // ensure our heartbeat has an extra chance so the time out is not too sensitive
	for {
		select {
		case _, ok := <-heartbeat: // always expect to receive a heartbeat.
			if !ok {
				return
			}
			fmt.Println("pulse")
		case r, ok := <-results:
			if !ok {
				return
			}
			fmt.Printf("results %v\n", r.Second())
		case <-time.After(timeout): // given no heartbeats or new results, time out.
			return
		}
	}
}
