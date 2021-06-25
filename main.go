package main

// 如果包名不是以 . 或 / 开头，如 "fmt" 或者 "container/list"，则 Go 会在全局文件进行查找；
// 如果包名以 ./ 开头，则 Go 会在相对目录中查找；
// 如果包名以 / 开头，则会在系统的绝对路径中查找。
import (
	"fmt"
)

// go run packages.go: 运行 Go 程序
// go run --work packages.go: 运行 Go 程序并打印出临时的工作目录
// go build packages.go: 生成可执行文件
func main() {
	fmt.Println("Go Practice!")
}
