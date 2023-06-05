package main

import (
	"fmt"
	"sync"
)

func main() {

	ch := make(chan string)
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		ls := []string{"Sarina Esmaeilzade", "Parsa Doosti", "Ehsan Alibaz"}
		for _, v := range ls {
			ch <- v
		}
	}()

	go func() {
		defer wg.Done()
		ls := []string{"Abolfazl Baeu", "Amir Norouzi", "Pedram Azarnoosh"}
		for _, v := range ls {
			ch <- v
		}
	}()

	go func() {
		wg.Wait()
		close(ch)
	}()

	for v := range ch {
		fmt.Println(v)
	}
}
