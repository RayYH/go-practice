package main

import (
	"fmt"
	"time"
)

func pump(ch chan int) {
	for i := 0; ; i++ {
		ch <- i
	}
}

func suck(ch chan int) {
	for {
		fmt.Println(<-ch)
	}
}

func main() {
	ch := make(chan int)
	go pump(ch)
	go suck(ch)
	time.Sleep(1e9)
}
