package main

import "fmt"

// A send operation on a channel (and the goroutine or
// function that contains it) blocks until a receiver is available

// A receive operation for a channel blocks (and the goroutine or
// function that contains it) until a sender is available

func pump(ch chan int) {
	for i := 0; ; i++ {
		ch <- i
	}
}

func main() {
	ch := make(chan int)
	go pump(ch)
	fmt.Println(<-ch) // only 0!! consumed only once!
}
