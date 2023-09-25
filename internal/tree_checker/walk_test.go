package treechecker_test

import (
	"fmt"
	"testing"
	"treechecker"

	"golang.org/x/tour/tree"
)

func TestWalk(t *testing.T) {
	tr := tree.New(1)
	ch := make(chan int)
	go treechecker.Walk(tr, ch)
	for i := 0; i < 10; i++ {
		value := <-ch
		if value != i+1 {
			fmt.Printf("expected %d, got %d\n", i, value)
			t.Fail()
		}
	}
}
