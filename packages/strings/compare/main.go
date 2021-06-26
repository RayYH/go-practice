package main

import (
	"fmt"
	"strings"
)

// Compare 用于比较字符串，返回 1、-1、0，该方法比 ==、>、< 操作要快
func main() {
	fmt.Println(strings.Compare("a", "b")) // -1
	fmt.Println(strings.Compare("a", "a")) // 0
	fmt.Println(strings.Compare("b", "a")) // 1
}
