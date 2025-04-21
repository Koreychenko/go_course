package main

import "fmt"

type List[T any] struct {
	next *List[T]
	val  T
}

func (l List[T]) GetAllNextValues() []T {
	var values []T

	values = append(values, l.val)

	if l.next != nil {
		values = append(values, l.next.GetAllNextValues()...)
	}

	return values
}

func prepareList[T comparable](values []T) []*List[T] {
	var nodes []*List[T]

	for i, value := range values {
		nodes = append(nodes, &List[T]{nil, value})
		if i > 0 {
			nodes[i-1].next = nodes[i]
		}
	}

	return nodes
}

func main() {
	nodes := prepareList([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	fmt.Println(nodes[1].GetAllNextValues())
}
