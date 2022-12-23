package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)

	go counter(ch)

	for i := range ch {
		fmt.Println("Data from channel______")
		fmt.Println(i)
	}
}

func counter(ch chan int) {
	for i := 0; i < 5; i++ {
		time.Sleep(5 * time.Second)
		ch <- i
	}

	close(ch)
}
