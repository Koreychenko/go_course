package main

import (
	"golang.org/x/tour/tree"
)

func walkHelper(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		walkHelper(t.Left, ch)
	}

	ch <- t.Value

	if t.Right != nil {
		walkHelper(t.Right, ch)
	}
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	walkHelper(t, ch)

	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go Walk(t1, ch1)
	go Walk(t2, ch2)

	for {
		val1, ok1 := <-ch1
		val2, ok2 := <-ch2

		if ok1 == ok2 && ok1 == false {
			break
		}

		if val1 != val2 {
			return false
		}

		continue
	}

	return true
}

func main() {
	print(Same(tree.New(1), tree.New(1)))
}
