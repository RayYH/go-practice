package main

import "fmt"

// fatal error: all goroutines are asleep - deadlock!
func main() {
	ch := make(chan int)
	// 由于没有消费者协程，下面的语句会阻塞，main 协程就会被阻塞住 (sleep)
	// 改为 go func() {ch <- 1}() 就可以了
	ch <- 1
	go func() {
		fmt.Println(<-ch)
	}()
}
