package main

import (
	"fmt"
	"pipelines"
	"treechecker"

	"golang.org/x/tour/tree"
)

func main() {
	// treeCheck()
	concur()
}

func treeCheck() {
	fmt.Println("Hello")
	t := tree.New(1)
	ch := make(chan int)
	go treechecker.Walk(t, ch)
	for i := range ch {
		fmt.Println(i)
	}
}

func concur() {
	pipelines.Pipeline([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
}
