package main

import "fmt"

func main() {
	// 你可以通过 &i 来获取变量 i 的内存地址
	a := 1
	fmt.Printf("%v\n", &a)

	// 除了 %v 打印描述符之外，%p 打印描述符也可以打印出指针的地址
	var i1 = 5
	fmt.Printf("An integer: %d, it's location in memory: %p\n", i1, &i1)

	// 可以直接使用 var 关键字声明一个指针类型的变量
	var intP *int
	// 在指针未被赋值前指针的值是 0，这时使用 *intP 访问会触发一个 panic
	fmt.Printf("The value of pointer is %p\n", intP)
	intP = &i1
	fmt.Printf("The value at memory location %p is %d\n", intP, *intP)
}
