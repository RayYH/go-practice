package main

import (
	"fmt"
	"time"
)

// fatal error: concurrent map read and map write
func main() {
	c := make(map[string]int)
	go func() {
		for j := 0; j < 1000000; j++ {
			c[fmt.Sprintf("%d", j)] = j
		}
	}()
	go func() {
		for j := 0; j < 1000000; j++ {
			fmt.Println(c[fmt.Sprintf("%d", j)])
		}
	}()

	time.Sleep(time.Second * 20)
}
