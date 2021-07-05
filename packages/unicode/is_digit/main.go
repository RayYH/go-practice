package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Println(unicode.IsDigit('1')) // true
	fmt.Println(unicode.IsDigit('a')) // false
	fmt.Println(unicode.IsDigit(1))   // false
}
