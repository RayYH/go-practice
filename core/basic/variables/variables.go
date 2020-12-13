package variables

import (
	"fmt"
	"math"
	"os"
)

// A var declaration creates a variable of a particular type, each declaration has the general form:
// var name type = expression
// Either the `type` or the `= expression` part may be omitted, but not both.
// If the expression is omitted, the initial value is the zero value for the type.
// So in Go, there is no uninitialized variable
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

// 函数可以返回多个具名返回值
func GetName() (firstName, middleName, lastName string) {
	// If an entity is declared within a function, it is local to that function.
	firstName = "Ray"
	middleName = "Young"
	lastName = "Hong"

	return
}

// If declared outside of a function, it's visible in all files of the package to which it belongs.
var declaredVariable float64

// The result list is omitted if the function doesn't return anything.
func init() {
	// 可以在 init 函数中对全局变量进行初始化
	declaredVariable = math.Atan(1)
}

// The case of the first letter of a name determines its visibility across package boundaries.
// If the name begins with a upper-case letter, it's exposed, which means it's visible and accessible
// outside of its own package
func IterateOverAString() {
	for index, s := range "Hello World" {
		fmt.Printf("The index number of %c is %d\n", s, index)
	}
}

// 可以在某个代码块的内层代码块中使用相同名称的变量，此时外部的同名变量将会暂时隐藏，你的所有操作都只会影响内部代码块的局部变量
