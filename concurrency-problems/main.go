package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	deadlock()

	starvation()

	race()
}

// deadlock simulates a deadlock scenario in which two concurrent tasks block each other and cause a deadlock
func deadlock() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	i, j := 0, 0

	go func() {
		for i == 0 {
			time.Sleep(time.Second)
		}

		j = 1
		fmt.Println("Done")
		wg.Done()
	}()

	go func() {
		for j == 0 {
			time.Sleep(time.Second)
		}

		i = 1
		fmt.Println("Done")
		wg.Done()
	}()

	wg.Wait()
}

// race simulates race condition in which value of i is indeterministic at the time of printing
func race() {
	for {
		i := 0

		go func() {
			fmt.Printf("i is sometimes: %d \n", i)
		}()

		go func() {
			i++
		}()
	}
}

// starvation simulates a greedy task which keeps shared resources busy
// and causes other tasks to starve
func starvation() {
	wg := sync.WaitGroup{}
	wg.Add(2)

	sharedResource := 0
	greedyTask := func() {
		if sharedResource == 0 {
			defer func() { sharedResource = 0 }()
			sharedResource = 1

			fmt.Println("I am greedy.")

			// Keep the shared resource busy
			time.Sleep(time.Millisecond * 100)
		}
	}

	starvingTask := func() {
		if sharedResource == 0 {
			defer func() { sharedResource = 0 }()
			sharedResource = 1

			fmt.Println("I am starving.")
		} else {
			time.Sleep(time.Second)
		}
	}

	go func() {
		for {
			greedyTask()
		}
	}()

	go func() {
		time.Sleep(time.Millisecond)
		for {
			starvingTask()
		}
	}()

	wg.Wait()
}
