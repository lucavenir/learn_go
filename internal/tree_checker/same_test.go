package treechecker_test

import (
	"fmt"
	"testing"
	"treechecker"

	"golang.org/x/tour/tree"
)

func TestSame(t *testing.T) {
	t1 := tree.New(1)
	t2 := tree.New(2)

	same := treechecker.Same(t1, t2)
	if same {
		fmt.Println("Two trees are different, but you say they've got the same Walk")
		t.Fail()
	}

	t1 = tree.New(20)
	t2 = tree.New(20)
	same = treechecker.Same(t1, t2)
	if !same {
		fmt.Println("Two trees have the same Walk, but you say they've got it different")
		t.Fail()
	}
}
