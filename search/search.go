package search

import (
	"cmp"
)

type Searching[T cmp.Ordered] []T

func (value Searching[T]) Len() int           { return len(value) }
func (value Searching[T]) Swap(i, j int)      { value[i], value[j] = value[j], value[i] }
func (value Searching[T]) Less(i, j int) bool { return value[i] < value[j] }

func (array *Searching[T]) BinarySearch(search T) (int, T) {
	i := 0
	j := len(*array)

	for i < j {
		mid := (i + j) / 2
		el := (*array)[mid]
		switch {
		case el < search:
			i = mid
		case el == search:
			return mid, el
		case el > search:
			j = mid
		}
	}

	panic("can't find 'search'")
}
