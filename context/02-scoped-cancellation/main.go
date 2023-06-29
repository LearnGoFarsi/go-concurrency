package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	StringProcessor(ctx)
	Vain(ctx)
}

func StringProcessor(ctx context.Context) {

	ctx, cancel := context.WithTimeout(ctx, time.Second*1)
	defer cancel()

	ch := make(chan string)
	wg := sync.WaitGroup{}
	wg.Add(1)

	writer := func(ctx context.Context) {
		defer wg.Done()
		ls := []string{"Zan", "Zendegi", "Azadi", "Mahsa", "Hameye dokhtarane iran", "Zan", "Zendegi", "Azadi", "Mahsa", "Hameye dokhtarane iran"}
		for _, s := range ls {
			select {
			case <-ctx.Done():
				return
			case ch <- s:
			}
		}
	}

	reader := func(ctx context.Context) {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				return
			case v, ok := <-ch:
				if !ok {
					return
				}
				fmt.Println(v)
			}
		}
	}

	go writer(ctx)

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go reader(ctx)
	}

	wg.Wait()

	fmt.Println("Processing finished")
}

func Vain(ctx context.Context) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()
	<-ctx.Done()
	fmt.Println("Vain finished")
}
