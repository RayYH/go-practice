package main

import "fmt"

// chan<- (channel can only receive data)
// <-chan (channel can only send data)

func producer(start, step, count int, out chan<- int) {
	for i := 0; i < count; i++ {
		out <- start
		start += step
	}
	close(out)
}

func consumer(in <-chan int, done chan<- bool) {
	for num := range in {
		fmt.Println(num)
	}
	done <- true
}

func main() {
	numChan := make(chan int)
	done := make(chan bool)
	go producer(0, 10, 10, numChan)
	go consumer(numChan, done)
	<-done
}
