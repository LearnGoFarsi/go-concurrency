package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(2)

	writer := func() {
		ls := []int{1, 2, 3, 4, 5, 6, 0}
		for _, s := range ls {
			ch <- s
		}

		close(ch)
		wg.Done()
	}

	reader := func() {
		for v := range ch {
			fmt.Println(v)
		}
		wg.Done()
	}

	go writer()
	go reader()

	wg.Wait()

	v, ok := <-ch
	if ok {
		fmt.Println("After closed:", v)
	}

}
