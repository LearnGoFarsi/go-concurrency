package main

import (
	"fmt"
	"time"
)

func main() {

	ch := make(chan int, 5)
	go func() {
		ch <- 1
		ch <- 2
		ch <- 3
		ch <- 4
		ch <- 5
		fmt.Println("I wrote 5 elements without waiting for the reader")
		ch <- 6 // this line is blocking till the reader starts to pick an element from the channel
		fmt.Println("I finished my work")
		close(ch)
	}()

	go func() {
		time.Sleep(time.Second)
		for i := range ch {
			fmt.Println(i)
		}
	}()

	time.Sleep(time.Second * 2)
}
