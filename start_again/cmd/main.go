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
}
