package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	// 根据指定条件将字符串切割成切片
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	fmt.Printf("Fields are: %q", strings.FieldsFunc("  foo1;bar2,baz3...", f))
}
