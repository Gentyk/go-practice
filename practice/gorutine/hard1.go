package main

import (
	"fmt"
	"sync"
	"context"
	"time"
)


func worker(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
			case <- ctx.Done():
				fmt.Println("good finish")
				return
			default:
				fmt.Println("working")
				time.Sleep(100*time.Millisecond)
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	// defer cancel()

	var wg sync.WaitGroup;

	wg.Add(2)
	go worker(ctx, &wg)
	go worker(ctx, &wg)

	time.Sleep(1 * time.Second)
    cancel() 
	fmt.Println("finish")

	wg.Wait()
	fmt.Println("finish2")
}