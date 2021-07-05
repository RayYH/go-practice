package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		// %b 是用于表示位的格式化标识符
		fmt.Printf("%b\n", i)
	}
}
