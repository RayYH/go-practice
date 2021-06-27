package functions

import "fmt"

func Greet() {
	fmt.Println("Hello")
}

func DeferOrders() {
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
}
