package main

// cgo 会替代 Go 编译器来产生可以组合在同一个包中的 Go 和 C 代码，如果你想要在你的 Go 程序中使用 cgo，则必须在单独的一行使用
// import "C" 来导入，一般来说你可能还需要 import "unsafe"。然后，你可以在 import "C" 之前使用注释的形式导入 C 语言库 (甚至有效的
// C 语言代码)，名称 "C" 并不属于标准库的一部分，这只是 cgo 集成的一个特殊名称用于引用 C 的命名空间。在这个命名空间里所包含的 C 类型都
// 可以被使用，例如 C.uint、C.long 等等，还有 libc 中的函数 C.random() 等也可以被调用。

/*
这段代码在 Windows 上不能运行，暂时 skip 掉
------------------------------------------------------------------------------------------------------------------------
import "C"


// #include <stdlib.h>
import "C"
import (
	"fmt"
	"time"
)

func Random() int {
	return int(C.random())
}

func Seed(i int) {
	C.srandom(C.uint(i))
}

func main() {
	now := time.Now().Unix()
	Seed(int(now))
	fmt.Println(Random())
}

------------------------------------------------------------------------------------------------------------------------
*/
