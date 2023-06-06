package main

import (
	"fmt"
	"io/ioutil"
	"sync"
)

func main() {

	files := make(chan int)
	dirs := make(chan string)
	tokens := make(chan struct{}, 10)

	wg := &sync.WaitGroup{}

	countFiles := func(dirPath string) {

		defer func() {
			wg.Done()
			<-tokens
		}()

		entries, err := ioutil.ReadDir(dirPath)
		if err != nil {
			return
		}

		for _, entry := range entries {
			if entry.IsDir() {
				dirs <- fmt.Sprintf("%s/%s", dirPath, entry.Name())
			} else {
				files <- 1
			}
		}
	}

	go func() {
		for dir := range dirs {
			wg.Add(1)
			tokens <- struct{}{}
			go func(dir string) {
				countFiles(dir)
			}(dir)
		}
	}()

	root := "."
	dirs <- root

	go func() {
		wg.Wait()
		close(files)
		close(dirs)
	}()

	cnt := 0
	for i := range files {
		cnt += i
	}

	fmt.Printf("Number of files in the %s directory and its subdirectories: %d \n", root, cnt)
}
