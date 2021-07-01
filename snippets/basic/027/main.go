package main

import "fmt"

type myStruct struct {
	i int
	f float32
	s string
}

func main() {
	ms := new(myStruct)
	ms.i = 1
	ms.f = float32(3.14)
	ms.s = "Ray"

	fmt.Printf("The int is %d\n", ms.i)
	fmt.Printf("The float32 is %f\n", ms.f)
	fmt.Printf("The string is %s\n", ms.s)
	fmt.Printf("%v", ms)
}
