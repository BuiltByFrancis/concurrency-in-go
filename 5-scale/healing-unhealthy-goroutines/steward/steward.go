package main

import (
	"log"
	"os"
	"time"
)

func main() {
	type startGoroutineFn func(
		done <-chan interface{},
		pulseInterval time.Duration,
	) (heartbeat <-chan interface{})

	newSteward := func(
		timeout time.Duration,
		startGoroutine startGoroutineFn,
	) startGoroutineFn {
		return func(
			done <-chan interface{},
			pulseInterval time.Duration,
		) <-chan interface{} {
			heartbeat := make(chan interface{})
			go func() {
				defer close(heartbeat)

				var wardDone chan interface{}
				var wardHeartbeat <-chan interface{}
				startWard := func() { // Consistent way of kicking of our ward go routine
					wardDone = make(chan interface{})                             // Always create a new done channel or context.
					wardHeartbeat = startGoroutine(or(wardDone, done), timeout/2) // Start the ward, done channel is a logical or of the steward or ward being done.
				}
				startWard()
				pulse := time.Tick(pulseInterval)

			monitorLoop:
				for {
					timeoutSignal := time.After(timeout)

					for { // Ensure the steward can send out its own pulse
						select {
						case <-pulse:
							select {
							case heartbeat <- struct{}{}:
							default:
							}
						case <-wardHeartbeat: // Monitor wards pulse.
							continue monitorLoop
						case <-timeoutSignal: // Given we time out, close the ward and restart it.
							log.Println("steward: ward unhealthy; restarting")
							close(wardDone)
							startWard()
							continue monitorLoop
						case <-done:
							return
						}
					}
				}
			}()

			return heartbeat
		}
	}
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ltime | log.LUTC)

	doWork := func(done <-chan interface{}, _ time.Duration) <-chan interface{} {
		log.Println("ward: Hello, I'm irresponsible!")
		go func() {
			<-done // Mock go routine does nothing, just waits to be cancelled. (No pulse is ever sent)
			log.Println("ward: I am halting.")
		}()
		return nil
	}
	doWorkWithSteward := newSteward(4*time.Second, doWork) // Create a steward for the doWork function (ward)

	done := make(chan interface{})
	time.AfterFunc(9*time.Second, func() { // End Example after 9 seconds
		log.Println("main: halting steward and ward.")
		close(done)
	})

	for range doWorkWithSteward(done, 4*time.Second) {
	} //Start the steward and range over its pulses to prevent the example from halting
	log.Println("Done")
}

func or(channels ...<-chan interface{}) <-chan interface{} {
	switch len(channels) {
	case 0:
		return nil
	case 1:
		return channels[0]
	}

	orDone := make(chan interface{})
	go func() {
		defer close(orDone)

		switch len(channels) {
		case 2:
			select {
			case <-channels[0]:
			case <-channels[1]:
			}
		default:
			select {
			case <-channels[0]:
			case <-channels[1]:
			case <-channels[2]:
			case <-or(append(channels[3:], orDone)...):
			}
		}
	}()
	return orDone
}
