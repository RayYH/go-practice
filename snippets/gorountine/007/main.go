package main

import (
	"fmt"
	"time"
)

func main() {
	// ch :=make(chan type, value)
	// value == 0 -> synchronous, unbuffered
	// value > 0 -> asynchronous, buffered
	ch := make(chan int, 10)
	go func() {
		time.Sleep(1 * time.Second)
		fmt.Println("received", <-ch)
	}()
	fmt.Println("sending", 10)
	ch <- 10
	fmt.Println("sent", 10)
	time.Sleep(2 * time.Second)
}
