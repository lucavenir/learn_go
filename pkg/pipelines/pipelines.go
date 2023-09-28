package pipelines

import (
	"fmt"
)

func Pipeline(input []int) {
	in := inputToChannel(input)
	squared := square(in)
	print(squared)
}

func inputToChannel(input []int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range input {
			out <- n
		}
		close(out)
	}()
	return out
}

func square(input <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range input {
			out <- n * n
		}
		close(out)
	}()
	return out
}
func print(input <-chan int) {
	for value := range input {
		fmt.Println(value)
	}
}
