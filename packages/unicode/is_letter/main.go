package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Println(unicode.IsLetter('a')) // true
	fmt.Println(unicode.IsLetter('1')) // false
	fmt.Println(unicode.IsLetter(1))   // false
	fmt.Println(unicode.IsLetter(' ')) // false
}
