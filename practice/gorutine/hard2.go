package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func worker(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("finish")
			return
		default:
			fmt.Println("work")
			time.Sleep(100 * time.Millisecond)
		}
	}

}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(2)
	go worker(ctx, &wg)
	go worker(ctx, &wg)

	wg.Wait()

}
