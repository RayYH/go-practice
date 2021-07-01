package main

import (
	"fmt"
	"math/rand"
)

func main() {
	p := rand.Intn(100)
	q := rand.Intn(100) + 101
	// 格式化输出时，%t 表示输出的值为布尔型
	fmt.Printf("%t", p == q)
}
