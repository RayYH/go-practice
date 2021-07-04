package main

import (
	// 可以将多个包别名为 .，相当于把包内的对象引入到当前包中，直接使用这些对象即可
	. "fmt"
	. "math"
)

func main() {
	Printf("Abs(-1) = %.2f", Abs(-1.23))
}
