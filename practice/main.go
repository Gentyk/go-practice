package main

import "fmt"

func main() {
	for i := 0; i < 3; i++ {
		fmt.Println(i)
	}

	for i := range 3 {
		fmt.Println(i)
	}
}
