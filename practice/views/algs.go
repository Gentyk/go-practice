package views

import (
	"unicode/utf8"
	// "fmt"
)

// разворот строки
func swapString(s string) string {
	n := utf8.RuneCountInString(s)
	if n < 2 {
		return s
	}
	runes := []rune(s)
	for i := range n / 2 {
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
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}

// сортировки
func swap(array []int, i, j int) {
	if array[i] > array[j] {
		array[i], array[j] = array[j], array[i]
	}
}

func bumbleSort(array []int) {
	n := len(array)
	for i := range n - 1 {
		for j := i + 1; j < n; j++ {
			swap(array, i, j)
		}
	}
}

func quickSort(a []int, start, end int) {
	// fmt.Println(a[start: end+1])
	if end-start < 0 {
		return
	}
	if end-start == 1 {
		swap(a, start, end)
		return
	}
	_start, _end := start, end
	med := end
	end -= 1
	for start < end {
		for a[start] < a[med] && start < end {
			start++
		}
		for a[end] >= a[med] && start < end {
			end--
		}
		if a[start] >= a[med] && a[end] < a[med] && start < end {
			a[start], a[end] = a[end], a[start]
		}
	}
	if a[start] >= a[med] {
		a[start], a[med] = a[med], a[start]
		med = start
	}
	// fmt.Println(a[_start: _end+1])
	quickSort(a, _start, med-1)
	quickSort(a, med+1, _end)
}
