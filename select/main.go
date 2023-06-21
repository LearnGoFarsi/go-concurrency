package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "Nima"
		// dont close the channel to keep the receiver blocking
	}()

	go func() {
		c2 <- "Shafigh doost"
		// dont close the channel to keep the receiver blocking
	}()

NIMA:
	for {
		select {
		case v := <-c1:
			fmt.Println(v)
		case v := <-c2:
			fmt.Println(v)
		case <-time.After(time.Second * 2):
			fmt.Println("Timeout")
			break NIMA
		}
	}
}
