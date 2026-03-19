package main

import "fmt"

func swap(s string) string {
	fmt.Println(s)
	tr := []rune(s)
	for i := 0; i < len(s)/2; i++ {
		tr[i], tr[len(s)-1-i] = tr[len(s)-1-i], tr[i]
	}
	s2 := string(tr)
	fmt.Println(s2)
	return s2
}

func delete(s string, t int) string {
	fmt.Println(s)
	tr := []rune(s)
	s2 := string(append(tr[:t], tr[t+1:]...))
	fmt.Println(s2)
	return s2
}

func main() {
	swap("abcde")
	delete("abcde", 1)
	swap("abcdef")
}
