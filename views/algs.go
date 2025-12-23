package views

import (
	"unicode/utf8"
)

// разворот строки
func swapString(s string) string {
	n := utf8.RuneCountInString(s)
	if n < 2 {
		return s
	}
	runes := []rune(s)
	for i := range n/2 {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}
	return string(runes)
}

// числа фиббоначи
func fib(n int) int {
	if n < 0 {
		return 0
	}
	if n < 3 {
		return n
	}
	return n * fib(n-1)
}

func fib2(n int) int {
	if n < 0 {
		return 0
	}
	if n < 2 {
		return n
	}
	result := 1
	for i := 2; i <= n; i ++ {
		result *= i
	}
	return result
}
