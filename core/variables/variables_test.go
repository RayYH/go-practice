package variables

import (
	"fmt"
	"reflect"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVariablesDeclarationAndInitialization(t *testing.T) {
	// 先声明类型、再进行赋值
	t.Run("declare first, then initialize", func(t *testing.T) {
		// var declares 1 or more variables
		var a, b int
		var c string
		a, b, c = 5, 7, "abc"
		assert.Equal(t, 5, a)
		assert.Equal(t, 7, b)
		assert.Equal(t, "abc", c)
	})

	// 可以直接使用 `var` 关键字声明变量并直接初始化，编译器会自动推断变量的类型
	t.Run("declaration and initialization at the same time", func(t *testing.T) {
		var v = 10 // the compiler will infer the type of variable `v`
		assert.Equal(t, v, 10)
	})

	// 使用 := 海象运算符声明变量并进行初始化、这种方法省去了 `var` 关键字
	// := 只能被用在函数体内，而不可以用于全局变量的声明与赋值
	t.Run(":= short assignment", func(t *testing.T) {
		d, e, f := 5, 7, "abc"
		assert.Equal(t, 5, d)
		assert.Equal(t, 7, e)
		assert.Equal(t, "abc", f)

		var g int
		// 已经定义过的变量也不能再使用 :=，但是在多个变量并行赋值时只要有一个变量未声明，我们就可以使用 := 海象运算符
		g, h, i := 1, 2, 3
		assert.Equal(t, 1, g)
		assert.Equal(t, 2, h)
		assert.Equal(t, 3, i)
	})

	// Go 中交换两个变量与 Python 中类似，直接对两个变量赋各自的新值即可
	t.Run("swap two variables' value", func(t *testing.T) {
		a := 5
		b := 7
		a, b = b, a
		assert.Equal(t, 7, a)
		assert.Equal(t, 5, b)
	})

	// 未进行初始化的变量会被赋予对应类型的零值
	t.Run("Variables declared without a corresponding initialization are zero-valued.", func(t *testing.T) {
		var v1 int
		var v2 float64
		var v3 string
		var v4 []int
		assert.Equal(t, v1, 0)
		assert.Equal(t, v2, 0.0)
		assert.Equal(t, v3, "")
		assert.Nil(t, v4)
	})

	// Go 中变量还可以是结构体、map、函数等类型
	t.Run("Advanced variables declaration", func(t *testing.T) {
		var v5 struct {
			f int
		} // struct
		var v6 struct{}        // empty struct
		var v7 *int            // pointer to int type
		var v8 map[string]int  // map
		var v9 func(a int) int // func
		assert.NotNil(t, v5)
		assert.NotNil(t, v6)
		assert.Zero(t, v5.f)
		assert.Nil(t, v7)
		assert.Nil(t, v8)
		assert.Nil(t, v9)
	})
}

func TestGlobalVariablesInitializationAndScope(t *testing.T) {
	// 全局变量遵从 Go 中可见性的约束，即大写字母开头可以跨包访问，小写字母开头只能在本包内访问
	// globalString 可以只对本包成员可见
	// GlobalString 可以被外部包访问
	t.Run("global variables are visible to other files", func(t *testing.T) {
		assert.Equal(t, "This is a string.", globalString)
		assert.Equal(t, "This is also a string.", GlobalString)
	})

	// 全局变量不同于常量，只要可见，即可以被修改
	t.Run("global variables can be modified", func(t *testing.T) {
		assert.Equal(t, 24, myAge)
		assert.Equal(t, "Ray", myName)
		myAge, myName = 25, "Ray Hong"
		assert.Equal(t, 25, myAge)
		assert.Equal(t, "Ray Hong", myName)
	})

	// 全局变量可以在 `init` 函数中初始化
	t.Run("global variables can be initialized inside init func", func(t *testing.T) {
		assert.Equal(t, 0.7853981633974483, declaredVariable)
	})
}

// 空白标识符 _ 表示匿名变量，可以理解为一个只写变量，主要用于丢弃函数的部分返回值
func TestUsingBlankIdentifierToDiscardValues(t *testing.T) {
	var _, age = "Ray", 24
	assert.Equal(t, 24, age)

	getName := func() (firstName, middleName, lastName string) {
		firstName, middleName, lastName = "Ray", "Young", "Hong"
		return
	}

	_, middle, _ := getName()
	assert.Equal(t, "Young", middle)
}

// Go 中仍然有指针 (引用) 的存在，与其他高级语言类似，Go 不支持对指针类型进行算术运算
// 我们可以通过指针来修改其所指变量的内容，这在某些时候很有用，比如传递数组指针可以直接修改原数组
func TestContentOfVariablesCanBeModifiedThroughPointers(t *testing.T) {
	var os = runtime.GOOS
	assert.NotNil(t, os)

	// &os 表示获取 os 的内存地址，q = p 表示将 p 的值 (os 的内存地址) 赋给了 q
	// q == p --> os (p and q are both pointers)
	// pointers holds the references to the variables
	var p = &os
	var q = p
	assert.Equal(t, p, q)
	assert.Equal(t, "*string", fmt.Sprint(reflect.TypeOf(p)))
	assert.Equal(t, *p, *q)
	assert.Equal(t, "string", fmt.Sprint(reflect.TypeOf(*q)))

	*p = "new string"
	assert.Equal(t, os, "new string")
	assert.Equal(t, p, q)
	assert.Equal(t, *p, *q)
}
