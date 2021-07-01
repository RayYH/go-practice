package variables

import (
	"math"
	"os"
)

// 在函数体外声明的变量即为全局变量，可以在整个包或者外部包 (被导出后) 使用，与所在的源文件无关。
// 在函数体内声明的变量称之为局部变量，它们的作用域只在函数体内，参数和返回值变量也是局部变量。
// 内部作用域的变量可以暂时覆盖 (隐藏) 外部作用域的同名变量，代内部代码块执行完毕 (内部同名变量被释放)，外部同名变量便可以恢复使用。

// globalString is visible within current package.
var globalString = "This is a string."

// GlobalString is also visible to other application which imported this package.
var GlobalString = "This is also a string."

// 当一个变量被声明之后，系统自动赋予它该类型的零值 (所有的内存在 Go 中都是经过初始化的)
// 全局变量允许只声明不使用
var emptyGlobalVar string

// 有的文章将这种写法称为『因式分解关键字』
// group of global variables.
var (
	myName = "Ray"
	myAge  = 24
)

// We can use some built-in functions when declaring global variables.
var (
	_ = os.Getenv("HOME")
	_ = os.Getenv("USER")
)

// No initialization but only declaration.
var declaredVariable float64

// init 函数不能被人为调用，该函数在每个包完成初始化后自动执行，并且执行优先级比 `main` 函数高
// 每个源文件都只能包含一个 init 函数
func init() {
	// 变量除了在全局声明中初始化，也可以在 init 函数中初始化
	declaredVariable = math.Atan(1)
}
