package main

import (
	"fmt"
	"time"
)

func suck(ch <-chan int) {
	for {
		fmt.Println(<-ch)
	}
}

func pump() chan int {
	ch := make(chan int)

	go func() {
		for i := 0; ; i++ {
			ch <- i
		}
	}()

	return ch
}

func main() {
	stream := pump()
	go suck(stream)
	time.Sleep(time.Millisecond)
}
