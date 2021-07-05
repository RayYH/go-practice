package main

import "fmt"

func main() {
	a := 1
	// &a returns the location of variable a
	fmt.Printf("%v\n", &a)

	var i1 = 5
	fmt.Printf("An integer: %d, it's location in memory: %p\n", i1, &i1)

	var intP *int
	// use *intP will trigger a panic error
	fmt.Printf("The value of pointer is %p\n", intP)
	intP = &i1
	fmt.Printf("The value at memory location %p is %d\n", intP, *intP)
}
