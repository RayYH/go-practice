package main

import "fmt"

func main() {
	str := "Go is a beautiful language!"
	fmt.Printf("The length of str is: %d\n", len(str))
	// ASCII 编码的字符占用 1 个字节，每个索引都指向不同的字符
	for ix := 0; ix < len(str); ix++ {
		fmt.Printf("Character on position %d is: %c \n", ix, str[ix])
	}

	str2 := "日本語"
	fmt.Printf("The length of str2 is: %d\n", len(str2))
	// 非 ASCII 编码的字符 (占有 2 到 4 个字节) 不能单纯地使用索引来判断是否为同一个字符
	for ix := 0; ix < len(str2); ix++ {
		fmt.Printf("Character on position %d is: %c \n", ix, str2[ix])
	}
}
