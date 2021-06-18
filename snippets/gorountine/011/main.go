package main

import (
	"fmt"
	"time"
)

func suck(ch <-chan int) {
	go func() {
		for v := range ch {
			fmt.Println(v)
		}
	}()
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
	suck(pump())
	time.Sleep(time.Millisecond)
}
