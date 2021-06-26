package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.ToValidUTF8("我 \xe2\xa1 中华", "❤️"))
}
