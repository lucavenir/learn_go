package main

import "fmt"

func main() {
	var p *int32
	fmt.Println(p)
	// fmt.Println(*p) THIS PANICS LOL

	a := &p
	fmt.Println(a)

	p = new(int32)
	fmt.Println(p)
	fmt.Println(*p)
	*p = 12
	fmt.Println(p)
	fmt.Println(*p)
}
