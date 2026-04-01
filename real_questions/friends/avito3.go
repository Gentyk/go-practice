// You can edit this code!
// Click here and start typing.
package friends

// Merge N каналов. Написать ф-цию func merge(cs ...chan int) (<- out chan int). Все значения должны перенаправляться в out канал.

import (
	"fmt"
	"sync"
)

func Merge[T any](channels ...<-chan T) <-chan T {
	result_chan := make(chan T)
	var wg sync.WaitGroup

	for _, current_chan := range channels {
		wg.Add(1)
		go func(chan1 <-chan T) {
			for i := range chan1 {
				result_chan <- i
			}
			wg.Done()
		}(current_chan)
	}

	go func() {
		wg.Wait()
		close(result_chan)
	}()
	return result_chan
}

func Merge2[T any](workersCount int, channels ...<-chan T) <-chan T {
	result_chan := make(chan T)
	semaphore := make(chan struct{}, workersCount)
	var wg sync.WaitGroup

	for _, c_chan := range channels {
		semaphore <- struct{}{}
		wg.Add(1)
		go func(c <-chan T) {
			defer func() {
				<-semaphore
				wg.Done()
			}()
			for i := range c {
				result_chan <- i
			}

		}(c_chan)
	}

	go func() {
		wg.Wait()
		close(result_chan)
	}()

	return result_chan
}

func Merge3[T any](workersCount int, channels ...<-chan T) <-chan T {
	result_chan := make(chan T)
	all_chans := make(chan (<-chan T), len(channels))
	var wg sync.WaitGroup

	for _, c := range channels {
		all_chans <- c
	}
	close(all_chans)

	wg.Add(workersCount)

	for i := 0; i < workersCount; i++ {
		go func() {
			defer wg.Done()
			for ch := range all_chans {
				for d := range ch {
					result_chan <- d
				}
			}

		}()
	}

	go func() {
		wg.Wait()
		close(result_chan)
	}()

	return result_chan
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		defer close(ch1)
		ch1 <- 1
		ch1 <- 2
	}()
	go func() {
		defer close(ch2)
		ch2 <- 3
		ch2 <- 4
	}()

	merged := Merge(ch1, ch2)
	for val := range merged { // Будет читать значения, пока канал не закроется
		fmt.Println(val) // Выведет 1, 2, 3, 4 в произвольном порядке
	}
	fmt.Println("Done")
	fmt.Println("Hello, 世界")
}
