package main

// несколько связанных горутин

import (
	"fmt"
	"sync"
)

func printer(chan_out <-chan int, wg *sync.WaitGroup) {
	for val := range chan_out {
		fmt.Println(val)
	}
	wg.Done()
}

func creater(chan_in chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i < 11; i++ {
		chan_in <- i
	}
	close(chan_in)
}

func merger(chan1 <-chan int, chan2 <-chan int, chan3 <-chan int, chan_out chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for chan1 != nil || chan2 != nil || chan3 != nil {
		select {
		case val, ok := <-chan1:
			if ok {
				chan_out <- val
			} else {
				chan1 = nil
			}
		case val, ok := <-chan2:
			if ok {
				chan_out <- val
			} else {
				chan2 = nil
			}
		case val, ok := <-chan3:
			if ok {
				chan_out <- val
			} else {
				chan3 = nil
			}

		}
	}
	close(chan_out)

}

func main() {
	chan1 := make(chan int)
	chan2 := make(chan int)
	chan3 := make(chan int)
	chan_out := make(chan int)
	var wg sync.WaitGroup

	wg.Add(5)
	go printer(chan_out, &wg)
	go merger(chan1, chan2, chan3, chan_out, &wg)
	go creater(chan1, &wg)
	go creater(chan2, &wg)
	go creater(chan3, &wg)
	wg.Wait()

}
