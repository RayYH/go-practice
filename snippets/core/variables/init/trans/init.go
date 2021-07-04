package trans

import "math"

// 你可以在使用 import 导入包之后定义或声明 0 个或多个常量、变量和类型
// 这些对象的作用域都是全局的 (在本包范围内)，所以可以被本包中所有的函数调用

var Pi float64

func init() {
	// Go 会先执行 init 函数，再执行 main 函数
	Pi = 4 * math.Atan(1)

	// init 函数也经常被用在当一个程序开始之前调用后台执行的 goroutine
}
