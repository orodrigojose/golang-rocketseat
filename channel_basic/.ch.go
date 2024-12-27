package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 3)

	go func() {
		for i := 0; i < 8; i++ {
			ch <- i
		}

		close(ch) // close and prevent deadlock
		fmt.Println("Escrita finalizada!")
	}()

	time.Sleep(1 * time.Second)

	// Esvaziar channel
	// <-ch

	for valor := range ch {
		fmt.Println("Leitura: ", valor)
	}
}
