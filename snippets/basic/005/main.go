package main

import (
	"fmt"
	// Go 不支持相对路径引入，你必须从 "$GOPATH" 环境变量中的 src 目录下的路径开始引入
	"github.com/rayyh/go-practice/snippets/basic/005/trans"
)

var twoPi = 2 * trans.Pi

func main() {
	fmt.Printf("2*Pi = %g\n", twoPi) // 2*Pi = 6.283185307179586
}
