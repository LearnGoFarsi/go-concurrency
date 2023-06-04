package main

import (
	"fmt"
	"time"
)

func main() {

	or := func(channels ...<-chan any) <-chan struct{} {

		closer := make(chan any)
		done := make(chan struct{})

		// To make sure the done channel will be closed only once
		go func() {
			<-closer
			close(done)
		}()

		for _, ch := range channels {
			go func(ch <-chan any) {
				select {
				//signal closer to close the done channel
				case x := <-ch:
					closer <- x
				// finish all the waiting goroutines
				case <-done:
					fmt.Println("Finished")
					return
				}
			}(ch)
		}

		return done
	}

	sig := func(ch <-chan time.Time) <-chan any {
		done := make(chan any)
		go func() {
			<-ch
			close(done)
		}()
		return done
	}

	start := time.Now()
	<-or(sig(time.After(time.Second)),
		sig(time.After(time.Second*2)),
		sig(time.After(time.Second*3)),
		sig(time.After(time.Second*4)),
		sig(time.After(time.Second*5)))

	fmt.Println(time.Now().Sub(start))
	time.Sleep(time.Second)
}
