package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	dbs := [...]int{1, 2, 3, 4, 5}
	t0 := time.Now()
	results := make([]int, 5, 5)
	ch := make(chan int)
	for _, db := range dbs {
		go muchDbSuchCallWow(int(db), ch)
	}
	for i := range dbs {
		result := <-ch
		results[i] = result
	}
	fmt.Printf("Time wasted: %v\n", time.Since(t0))
	fmt.Println("Result:", results)
}

func muchDbSuchCallWow(id int, ch chan int) {
	// mocking of a db connection
	delay := rand.Float32() * 2000
	time.Sleep(time.Duration(delay) * time.Millisecond)
	ch <- int(id) * int(delay)
	fmt.Printf("Much db, such call, wowowow: %d\n", id)
}
