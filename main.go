package main

import (
	"fmt"
	"treechecker"

	"golang.org/x/tour/tree"
)

func main() {
	fmt.Println("Hello")
	t := tree.New(1)
	ch := make(chan int)
	go treechecker.Walk(t, ch)
	for i := range ch {
		fmt.Println(i)
	}
}
