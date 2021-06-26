package main

import (
	"fmt"
	"strings"
)

// Contains 用于确认字符串是否包含一个子串
func main() {
	fmt.Println(strings.Contains("seafood", "foo")) // true
	fmt.Println(strings.Contains("seafood", "bar")) // false
	fmt.Println(strings.Contains("seafood", ""))    // true
	fmt.Println(strings.Contains("", ""))           // true
}
