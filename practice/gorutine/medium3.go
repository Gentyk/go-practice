package main

import (
	"fmt"
	"sync"
	"math/rand"
	"math"
)

import (
	"fmt"
	"math/rand"
	"sync"
)

func generate(ch_out chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(ch_out)
	fmt.Println("start")
	for i := 0; i < 50; i++ {
		ch_out <- rand.Int()
		fmt.Println("added value")
	}
}

func x2(chan_in chan int, chan_out chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(chan_out)
	for value := range chan_in {
		chan_out <- value * value
		fmt.Println("up value")
	}
}

func filter1(ch_in chan int, ch_out chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(ch_out)

	for i := range ch_in {
		if i%2 == 0 {
			ch_out <- i
			fmt.Println("set value")
		}
	}
}

func printer(ch_in chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := range ch_in {
		fmt.Println("value", i)
	}
}

func main() {
	var wg sync.WaitGroup

	wg.Add(4)
	chan1 := make(chan int)
	chan2 := make(chan int)
	chan3 := make(chan int)

	go generate(chan1, &wg)
	go x2(chan1, chan2, &wg)
	go filter1(chan2, chan3, &wg)
	go printer(chan3, &wg)

	wg.Wait()
}
