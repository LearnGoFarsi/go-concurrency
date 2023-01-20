package main

import (
	"sync"
)

func main() {

	// factorial(10) = 55
	if sequential_factorial(10) != concurrent_factorial(10) {
		panic("the calculated factorial is incorrect.")
	}
}

func sequential_factorial(f int) int {
	sum := 0
	for i := 1; i < f+1; i++ {
		sum += i
	}

	return sum
}

func concurrent_factorial(f int) int {

	// using waitgroup we can wait for execution of a group of goroutines.
	wg := sync.WaitGroup{}
	wg.Add(f)

	// A Mutex is a mutual exclusion lock that can be used to synchronize access to a shared resource
	mux := sync.Mutex{}

	// shared memory location/address (shared resource)
	sum := 0

	for i := 1; i < f+1; i++ {
		// Each goroutine tries to access a shared resource
		go func(j int) {
			// begin of critical section
			// using mutex we synchronize access to the shared resource
			mux.Lock()
			sum += j
			mux.Unlock()
			// end of critical section

			wg.Done()
		}(i)
	}

	// wait for all the goroutines registered in the waitgroup to finish their work
	wg.Wait()

	return sum
}
