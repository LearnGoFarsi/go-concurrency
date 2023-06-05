package main

import "fmt"

func main() {
	readOnlyCh := Generator()
	for v := range readOnlyCh {
		fmt.Println(v)
	}
}

func Generator() <-chan string {
	c := make(chan string)
	go func() {
		defer close(c)
		ls := []string{"Setareh Tajik", "Nika Shakarami", "Zakaria Khial", "Mohammad reza Sarvi"}
		for _, v := range ls {
			c <- v
		}
	}()
	return c
}
