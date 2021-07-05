package main

import "fmt"

func main() {
	i := 42
	j := 3.14
	k := 0.867 + 0.5i

	// %T 用于打印出值的类型
	fmt.Printf("%T %T %T", i, j, k)
}
