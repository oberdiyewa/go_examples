package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	func() {
		fmt.Println("started")
	}()

	go counter(ch)

	for i := range ch {
		fmt.Println("Data from channel______")
		fmt.Println(i)
	}
}

func counter(ch chan int) {
	for i := 0; i < 50; i++ {
		time.Sleep(1000 * time.Millisecond)
		ch <- i
	}

	close(ch)
}
