package main

import (
	"fmt"
	"sync"
	"time"
)
func worker(ch chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	defer func() { <-ch }()

	fmt.Println("in progress")
	time.Sleep(50 * time.Millisecond)
}

func main() {
	workerCount := 1000
	onMomentCount := 5

	var wg sync.WaitGroup

	ch := make(chan struct{}, onMomentCount)

	for i := 0; i < workerCount; i++ {
		ch <- struct{}{}
		wg.Add(1)
		go worker(ch, &wg)
	}

	wg.Wait()
	fmt.Println("finish")
}
