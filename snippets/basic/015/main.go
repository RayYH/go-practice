package main

import "fmt"

// Rope 是 string 的别名类型
type Rope string

func main() {
	var a Rope
	a = "Hello"
	fmt.Println(a)
}
