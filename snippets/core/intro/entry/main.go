package main

import "fmt"

// 在 Go 中，每一个可执行程序都必须包含 main 函数，main 函数既不允许有参数，也不允许有返回类型
func main() {
	// 不要使用内置的 print、println 等全局打印函数 (在 Go 未来的发布版本中，这些函数可能会被移除掉)
	// 你应当始终使用 fmt 包提供的打印函数
	fmt.Println("Hello World")
}
