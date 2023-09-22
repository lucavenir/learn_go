package main

import (
	"fmt"
)

func main() {
	a := []rune("Amioèe€")
	b := '@'
	fmt.Println(a)
	fmt.Println(string(a))
	fmt.Println(b)
	fmt.Println(string(b))
}
