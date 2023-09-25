package treechecker

import (
	"golang.org/x/tour/tree"
)

func Same(t1 *tree.Tree, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	go Walk(t1, ch1)

	ch2 := make(chan int)
	go Walk(t2, ch2)

	for i1 := range ch1 {
		i2 := <-ch2
		if i1 != i2 {
			return false
		}
	}

	return true
}
