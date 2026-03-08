package main
import (
	"fmt"
	"sync"
)

func main() {
	chan1 := make(chan int, 10)

	var s sync.WaitGroup
	s.Add(2)

	go func() {
		for i := 1; i < 11; i++ {
			chan1 <- i
		}
		close(chan1)
		s.Done()
	}()

	go func() {
		for data := range chan1 {
			fmt.Println(data)
		}
		s.Done()
	}()

	s.Wait()

}