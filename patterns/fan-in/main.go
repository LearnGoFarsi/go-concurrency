package main

import (
	"fmt"
	"sync"
)

func main() {

	gen := func() <-chan int {
		ch := make(chan int)
		go func() {
			for i := 0; i <= 20; i++ {
				ch <- i
			}
			close(ch)
		}()
		return ch
	}

	ch1 := gen()
	ch2 := gen()
	ch3 := gen()
	ch4 := gen()

	fanin := func(chs ...<-chan int) <-chan int {
		out := make(chan int)

		go func() {
			wg := sync.WaitGroup{}
			wg.Add(len(chs))

			for _, ch := range chs {
				go func(ch <-chan int) {
					for i := range ch {
						out <- i
					}
					wg.Done()
				}(ch)
			}

			wg.Wait()
			close(out)
		}()

		return out
	}

	result := fanin(ch1, ch2, ch3, ch4)
	for i := range result {
		fmt.Println(i)
	}

}
