package search

import (
	"sort"
	"testing"
)

func TestBinary(t *testing.T) {
	var array Searching[int] = []int{-1, -0, 1, 16, 2, 3, 45, 4, 5, 6, -99, 7}

	sort.Sort(array)

	for i, el := range array {
		index, match := array.BinarySearch(el)
		if i != index || el != match {
			t.Fatalf("Didn't match %d", el)
		}
	}
}
