package main

import (
	"fmt"
	"sync"
)

func main() {

	ch := make(chan string)
	wg := sync.WaitGroup{}
	wg.Add(1)

	tasks := func() {
		ls := []string{"Zan", "Zendegi", "Azadi", "Mahsa", "Hameye dokhtarane iran", "Zan", "Zendegi", "Azadi", "Mahsa", "Hameye dokhtarane iran"}
		for _, s := range ls {
			ch <- s
		}

		close(ch)
		wg.Done()
	}

	workers := func() {
		for v := range ch {
			fmt.Println(v)
		}
		wg.Done()
	}

	go tasks()

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go workers()
	}

	wg.Wait()
}
