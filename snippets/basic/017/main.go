package main

import (
	"fmt"
	"unicode"
)

func main() {
	var ch1, ch2, ch3 rune
	ch1, ch2, ch3 = '1', 'A', ' ' // 注意这里的 '1' 如果改为 1 则输出 false
	fmt.Println(unicode.IsDigit(ch1))
	fmt.Println(unicode.IsLetter(ch2))
	fmt.Println(unicode.IsSpace(ch3))
}
