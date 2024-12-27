package main

import (
	"fmt"
	"time"
)

func producer(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
	}

	close(ch)
}

func consumer(ch chan int) {
	for v := range ch {
		fmt.Println(v)
	}

	fmt.Println("Consumer finalizado!")
}

func main() {
	ch := make(chan int)

	go producer(ch)
	go consumer(ch)
	go consumer(ch)

	time.Sleep(time.Second * 1)
}
