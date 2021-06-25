package main

import (
	. "fmt" // 使用 . 时无需加上包前缀，但为了避免命名空间污染和歧义，不建议这么使用
	. "math"
)

func main() {
	Println("Hello World", Abs(-1))
}
