package main

import "fmt"

func main() {
	ch := make(chan int)
	go func(x, y int, ch chan int) {
		ch <- x + y
	}(5, 6, ch)
	fmt.Println(<-ch)
}
