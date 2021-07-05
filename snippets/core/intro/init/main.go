package main

import "fmt"

func main() {
	fmt.Println("main")
}

func init() {
	fmt.Println("init 1")
}

func init() {
	fmt.Println("init 2")
}
