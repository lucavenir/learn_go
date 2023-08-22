package main

import "fmt"

func main() {
	i, j := 42, 2701.0

	p := &i
	fmt.Println(*p)
	fmt.Println(p)

	*p = 21
	fmt.Println(i)

	a := &j
	*a = *a / 37.0
	fmt.Println(j)

	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
