package main

import "fmt"

// 在使用格式化说明符时，可以使用 %v 来表示复数，但当你希望只表示其中的一个部分的时候需要使用 %f
func main() {
	a := 1 + 2i
	fmt.Printf("%v\n", a)
	fmt.Printf("%f\n", a)
	fmt.Printf("%v\n", real(a))
	fmt.Printf("%f\n", real(a))
	fmt.Printf("%v\n", imag(a))
	fmt.Printf("%f\n", imag(a))
}
