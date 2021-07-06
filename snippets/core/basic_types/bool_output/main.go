package main

import (
	"fmt"
	"math/rand"
)

func main() {
	p := rand.Intn(100)
	q := rand.Intn(100) + 101
	fmt.Printf("%t", p == q)
}
