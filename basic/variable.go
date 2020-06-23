package basic

import (
	"math"
	"os"
)

// 声明变量的一般形式是使用 var 关键字: var identifier [type] = value
// 但大多数时候 (在对一个变量声明并进行赋值的时候)，我们会选择省略类型 (因为编译器会自动推断其类型) var identifier = value
// 无需显式声明其为 string，因为 Go 编译器会自动推断其类型
var globalString = "This is a string."

// 全局变量可以并行赋值 (因式分解关键字)
var (
	myName = "Ray"
	myAge  = 24
)

// 声明全局变量时可以调用标准库函数
var (
	HOME = os.Getenv("HOME")
	USER = os.Getenv("USER")
)

var declaredVariable float64

func init() {
	// 可以在 init 函数中对全局变量进行初始化
	declaredVariable = math.Atan(1)
}

// 当一个变量被声明之后，系统自动赋予它该类型的零值:
// int 为 0，float 为 0.0，bool 为 false，string 为空字符串，指针为 nil
// 因此所有的内存在 Go 中都是经过初始化的

// 1. 如果你的全局变量希望能够被外部包所使用，则需要将首个单词的首字母大写
// 2. 可以在某个代码块的内层代码块中使用相同名称的变量，此时外部的同名变量将会暂时隐藏，你的所有操作都只会影响内部代码块的局部变量
