package search

import (
	"sort"
	"testing"
)

func TestBinary(t *testing.T) {
	var array Searching[int] = []int{-1, -0, 1, 16, 2, 3, 45, 4, 5, 6, -99, 7}

	sort.Sort(array)

	for _, el := range array {
		_, _, err := array.BinarySearch(el)
		if err != nil {
			t.Fatalf("Didn't match %d", el)
		}
	}
}
