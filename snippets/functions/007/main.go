package main

import "fmt"

const LIMIT = 40

var fibs [LIMIT + 1]uint64

func init() {
	fibs[0] = 0
	fibs[1] = 1
	for i := 2; i <= LIMIT; i++ {
		fibs[i] = fibs[i-1] + fibs[i-2]
	}
}

func main() {
	for i := 0; i <= LIMIT; i++ {
		fmt.Printf("fibs[%d] = %d\n", i, fibs[i])
	}
}
