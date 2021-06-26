package main

import (
	"fmt"
	"strings"
)

func main() {
	// 根据空白字符将字符串切割成切片
	fmt.Printf("Fields are: %q", strings.Fields("  foo bar  baz   "))
}
