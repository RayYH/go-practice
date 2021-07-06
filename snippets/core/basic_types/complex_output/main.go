package main

import "fmt"

func main() {
	a := 1 + 2i
	fmt.Printf("%v\n", a)
	fmt.Printf("%f\n", a)
	fmt.Printf("%v\n", real(a))
	fmt.Printf("%f\n", real(a))
	fmt.Printf("%v\n", imag(a))
	fmt.Printf("%f\n", imag(a))
}
