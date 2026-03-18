package main

// написать генератор, создающий слайс интов произвольной длины

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	for _, i := range genIntNew(10) {
		fmt.Println(i)
	}
}

func genIntNew(n int) []int {
	rand.Seed(time.Now().UnixNano())
	result := make([]int, 0, n)
	keys := make(map[int]struct{})
	i := 0
	for i < n {
		t := rand.Int()
		if _, ok := keys[t]; ok {
			continue
		}
		keys[t] = struct{}{}
		result = append(result, t)
		i++
	}
	return result
}
