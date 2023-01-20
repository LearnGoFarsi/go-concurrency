package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	wg := sync.WaitGroup{}

	// A RWMutex is a reader/writer mutual exclusion lock. The lock can be held by an arbitrary number of readers or a single writer.
	mux := sync.RWMutex{}

	sharedValue := 1

	read := func() {
		start := time.Now()

		mux.RLock()
		time.Sleep(time.Millisecond * 100)
		mux.RUnlock()

		fmt.Println(fmt.Printf("It took me %d milliseconds to read.", time.Now().Sub(start).Milliseconds()))
		wg.Done()
	}

	write := func() {

		start := time.Now()
		for mux.TryLock() == false {
			time.Sleep(time.Millisecond * 10)
		}

		sharedValue++
		mux.Unlock()

		fmt.Println(fmt.Printf("It took me %d milliseconds to write.", time.Now().Sub(start).Milliseconds()))

		wg.Done()
	}

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go read()
	}

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go write()
	}

	wg.Wait()
}
