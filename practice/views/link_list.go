package views

import "fmt"

// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}

func NewList[T any](v T) *List[T] {
	return &List[T]{val: v}
}

func (l *List[T]) Pretend(v T) *List[T] {
	return &List[T]{val: v, next: l}
}

func (l *List[T]) lenList() int {
	if l.next == nil {
		return 1
	}
	return 1 + l.next.lenList()
}

func run() {
	s := []int{12, 23, 33}
	newList := NewList(1)
	for _, v := range s {
		newList = newList.Pretend(v)
	}
	fmt.Println(s)
	fmt.Println(newList)
	fmt.Println(newList.lenList())
}
