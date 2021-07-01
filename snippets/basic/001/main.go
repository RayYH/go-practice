package main

import "fmt"

// main 函数是每一个可执行程序所必须包含的，该函数既没有参数，也没有返回类型
func main() {
	// 不要使用内置的 print、println 等全局打印函数，你应当始终使用 fmt 包提供的打印函数
	fmt.Println("Hello World")
}
