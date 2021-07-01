package main

import (
	"fmt"
	"os"
	"strconv"
)

// Go 语言的函数经常使用两个返回值来表示执行是否成功，我们通常使用 v, e := func(...)，这种形式称之为 comma,ok 模式
func main() {
	var orig = "ABC"
	var newS string

	fmt.Printf("The size of ints is: %d\n", strconv.IntSize)

	an, err := strconv.Atoi(orig)

	if err != nil {
		fmt.Printf("orig %s is not an integer - exiting with error\n", orig)
		os.Exit(1)
	}

	fmt.Printf("The integer is %d\n", an)
	an += 5
	newS = strconv.Itoa(an)
	fmt.Printf("The new string si %s\n", newS)
}
