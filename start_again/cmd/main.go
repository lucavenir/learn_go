package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello world!")
	var i uint64
	fmt.Println(i)

	switch i {
	case 0:
		backTicks := `Hello world`
		fmt.Println(backTicks)
	}
	arr := [...]uint16{1, 2, 3}
	fmt.Println(arr)
	slice := []uint16{1, 2, 3}
	anotherSlice := append(slice, 0)
	fmt.Println(slice)
	fmt.Println(anotherSlice)
	spread := make([]uint16, len(slice), 2*cap(slice))
	copy(spread, slice)
	fmt.Println(spread)
	spread = append(spread, anotherSlice...)
	fmt.Println(spread)
}
