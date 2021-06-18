package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go func() {
		time.Sleep(15 * time.Second)
		fmt.Println("received", <-ch)
	}()
	fmt.Println("sending", 10)
	ch <- 10
	fmt.Println("sent", 10)
}
