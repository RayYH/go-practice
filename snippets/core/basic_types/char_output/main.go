package main

import "fmt"

func main() {
	var ch int = '\u0041'
	var ch2 int = '\u03B2'
	var ch3 int = '\U00101234'

	// 格式化说明符 %c 用于表示字符
	// %v 或 %d 会输出用于表示该字符的整数
	// %X 按十六进制输出 UTF-8 数值
	// %U 输出格式为 U+hhhh 的字符串

	fmt.Printf("%d - %d - %d\n", ch, ch2, ch3) // integer
	fmt.Printf("%c - %c - %c\n", ch, ch2, ch3) // character
	fmt.Printf("%X - %X - %X\n", ch, ch2, ch3) // UTF-8 bytes
	fmt.Printf("%U - %U - %U", ch, ch2, ch3)   // UTF-8 code point
}
