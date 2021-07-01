package main

import "fmt"

func main() {
	// 你可以通过 &i 来获取变量 i 的内存地址
	a := 1
	fmt.Printf("%v\n", &a)

	var i1 = 5
	fmt.Printf("An integer: %d, it's location in memory: %p\n", i1, &i1)

	var intP *int
	intP = &i1
	fmt.Printf("The value at memory location %p is %d\n", intP, *intP)
}
