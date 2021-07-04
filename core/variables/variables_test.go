package variables

import (
	"fmt"
	"reflect"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeclarationAndInitialization(t *testing.T) {
	// 先声明变量的类型 (这时候变量已经被初始化为对应类型的零值)、再对变量进行赋值
	t.Run("declare first, then initialize", func(t *testing.T) {
		var a, b int
		var c string
		a, b, c = 5, 7, "abc"
		assert.Equal(t, 5, a)
		assert.Equal(t, 7, b)
		assert.Equal(t, "abc", c)
	})

	// 直接使用 var 关键字对变量进行声明和初始化，由于编译器会自动推断出变量的类型，因此你不必显式指出变量的类型
	t.Run("declaration and initialization at the same time", func(t *testing.T) {
		var v = 10
		assert.Equal(t, 10, v)
	})

	// 1. := 运算符表示对变量进行声明和初始化，此时变量的类型将由编译器自动推断
	// 2. := 只能被用在函数体中，不能用于全局变量的声明与赋值
	// 3. := 只能用于声明并初始化还未声明的变量，对于已经声明过的变量不能再使用 :=
	t.Run(":= short assignment", func(t *testing.T) {
		d, e, f := 5, 7, "abc"
		assert.Equal(t, 5, d)
		assert.Equal(t, 7, e)
		assert.Equal(t, "abc", f)

		var g, h int
		// 多个变量并行赋值时只要有一个变量未声明 (如下例中的 i)，我们就可以使用 := 海象运算符
		g, h, i := 1, 2, 3
		assert.Equal(t, 1, g)
		assert.Equal(t, 2, h)
		assert.Equal(t, 3, i)
	})

	// Go 中交换两个变量与 Python 中类似：a, b = b, a
	t.Run("swap two variables' value", func(t *testing.T) {
		a := 5
		b := 7
		a, b = b, a
		assert.Equal(t, 7, a)
		assert.Equal(t, 5, b)
	})

	// 只进行声明但未进行初始化的变量会被赋予对应类型的零值
	t.Run("variables declared without a corresponding initialization are zero-valued", func(t *testing.T) {
		var v1 int
		var v2 float64
		var v3 string
		assert.Equal(t, 0, v1)
		assert.Equal(t, 0.0, v2)
		assert.Equal(t, "", v3)
	})

	// Go 中变量还可以是切片、结构体、map、函数等多种类型，对于引用类型，变量声明后的零值都是 Nil
	// 由于结构体是值类型，因此结构体类型的变量对应的零值不是 Nil
	t.Run("declaration of variables holding advance types", func(t *testing.T) {
		var v4 []int
		var v5 struct { // struct
			f int
		}
		var v6 struct{}        // empty struct
		var v7 *int            // pointer to int type
		var v8 map[string]int  // map
		var v9 func(a int) int // func
		assert.Nil(t, v4)
		assert.NotNil(t, v5)
		assert.NotNil(t, v6)
		assert.Zero(t, v5.f)
		assert.Nil(t, v7)
		assert.Nil(t, v8)
		assert.Nil(t, v9)
	})
}

func TestGlobalVariablesInitializationAndScope(t *testing.T) {
	// 全局变量遵从 Go 中可见性的约束，以大写字母开头的变量可以跨包访问，以小写字母开头的变量只能在本包内访问
	t.Run("global variables are visible to other files", func(t *testing.T) {
		assert.Equal(t, "This is a string.", globalString)
		assert.Equal(t, "This is also a string.", GlobalString)
	})

	// 只进行声明，但是没有显式赋值的全局变量拥有默认的零值
	t.Run("global variables declared but no initialization have zero values", func(t *testing.T) {
		assert.Equal(t, "", emptyGlobalVar)
	})

	// 全局变量不同于常量，只要可见，就可以被修改
	t.Run("global variables can be modified", func(t *testing.T) {
		assert.Equal(t, 24, myAge)
		assert.Equal(t, "Ray", myName)
		myAge, myName = 25, "Ray Hong"
		assert.Equal(t, 25, myAge)
		assert.Equal(t, "Ray Hong", myName)
	})

	// 可以在 init 函数中对全局变量进行初始化
	t.Run("global variables can be initialized inside init func", func(t *testing.T) {
		assert.Equal(t, 0.7853981633974483, declaredVariable)
	})
}

// 空白标识符 _ 表示匿名变量，可以理解为一个只写变量，主要用途是丢弃函数的部分返回值
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

// 1. Go 中仍然有指针 (引用) 的存在，与其他高级语言类似，Go 并不支持对指针类型的变量进行算术运算
// 2. 值类型的变量的值存储在栈中，引用类型所指的变量会存储在堆中，以便进行垃圾回收，且比栈拥有更大的内存空间
// 3. 我们可以通过指针来修改其所指变量的内容，比如传递数组指针可以直接修改原数组
// 4. slice、map、channel 都是引用类型
func TestContentOfVariablesCanBeModifiedThroughPointers(t *testing.T) {
	var os = runtime.GOOS
	assert.NotNil(t, os)

	// Go 语言的取地址符是 &，放到一个变量前使用就会返回相应变量的内存地址
	var p = &os
	// 将一个引用类型赋给另一个变量时，只有引用 (地址) 被复制
	var q = p

	// p 和 q 都是 *string 类型 (指向字符串的指针)
	assert.Equal(t, p, q)
	assert.Equal(t, "*string", fmt.Sprint(reflect.TypeOf(p)))
	// *p 和 *q 都是字符串类型
	assert.Equal(t, *p, *q)
	assert.Equal(t, "string", fmt.Sprint(reflect.TypeOf(*q)))

	// 如果一个变量的值被修改了，那么这个值的所有引用都会指向被修改后的内容
	*p = "new string"
	assert.Equal(t, os, "new string")
	assert.Equal(t, p, q)
	assert.Equal(t, *p, *q)
}
