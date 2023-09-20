package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string, 10)
	for i := 0; i < 15; i++ {
		go sayWithChan(ch)
		ch <- fmt.Sprint(i)
	}
	fmt.Println("Heyyyyyyy")
	for s := range ch {
		fmt.Println(s) // deadlock
	}
}

func say(s string) {
	time.Sleep(2 * time.Second)
	fmt.Println(s)
}

func sayWithChan(ch chan string) {
	data := <-ch
	time.Sleep(2 * time.Second)
	fmt.Println(data)
	ch <- data
}
