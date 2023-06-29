package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {

	ctx := context.Background()

	writer := func(ctx context.Context, wg *sync.WaitGroup) (<-chan string, <-chan struct{}) {

		dataCh := make(chan string)
		signalCh := make(chan struct{})

		go func() {
			defer close(dataCh)
			defer close(signalCh)
			defer wg.Done()

			ls := []string{"Zan", "Zendegi", "Azadi", "Mahsa", "Hameye dokhtarane iran", "Zan", "Zendegi", "Azadi", "Mahsa", "Hameye dokhtarane iran", "Zan", "Zendegi", "Azadi", "Mahsa", "Hameye dokhtarane iran", "Zan", "Zendegi", "Azadi", "Mahsa", "Hameye dokhtarane iran", "Zan", "Zendegi", "Azadi", "Mahsa", "Hameye dokhtarane iran", "Zan", "Zendegi", "Azadi", "Mahsa", "Hameye dokhtarane iran"}

			for i := 0; i < len(ls); {

				var timeout time.Duration
				if i < 5 {
					timeout = time.Millisecond * 10
				} else if i < 10 {
					timeout = time.Millisecond * 20
				} else {
					timeout = time.Millisecond * 40
				}

				select {
				case <-ctx.Done():
					return
				case <-time.After(timeout):
					signalCh <- struct{}{}
				case dataCh <- ls[i]:
					i++
				}
			}
		}()

		return dataCh, signalCh
	}

	reader := func(ctx context.Context, dataCh <-chan string, wg *sync.WaitGroup, timeout time.Duration) {
		defer wg.Done()
		defer func() {
			fmt.Println("***Bye")
		}()

		for {
			time.Sleep(time.Millisecond * 100)

			select {
			case <-ctx.Done():
				return
			case v, ok := <-dataCh:
				if !ok {
					return
				}
				fmt.Println(v)
			case <-time.After(timeout):
				return
			}
		}
	}

	wg := sync.WaitGroup{}
	wg.Add(1)
	dataCh, signalCh := writer(ctx, &wg)

	for range signalCh {
		wg.Add(1)
		fmt.Println("***New reader added")
		go reader(ctx, dataCh, &wg, time.Millisecond*100)
	}

	wg.Wait()
}
