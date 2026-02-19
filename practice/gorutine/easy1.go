package main

import (
	"sync"
	"fmt"
)

func main() {

	var group1 sync.WaitGroup

	// group1.Add(5) - опасно выделять заранее, тк кто-то потом может изменить цикл
	for i:=0; i<5; i++ {
		group1.Add(1)
		go func(i int) {
			fmt.Println(i)
			defer group1.Done() 
		}(i)
	}

	group1.Wait()
}
