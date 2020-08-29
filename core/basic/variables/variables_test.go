package variables

import (
	"github.com/stretchr/testify/assert"
	"runtime"
	"strings"
	"testing"
)

// 全局变量可以在函数内使用
func TestGlobalString(t *testing.T) {
	assert.Equal(t, "This is a string.", globalString)
	assert.Equal(t, 24, myAge)
	assert.Equal(t, "Ray", myName)
}

func TestHomeAndUserNotEmpty(t *testing.T) {
	if len(HOME) == 0 {
		t.Error("HOME is empty")
	}

	if len(USER) == 0 {
		t.Error("USER is empty")
	}
	// 声明变量可以声明类型再赋值，也可以声明类型并同时赋值
	// 局部变量赋值可以直接使用 :=
	// 下面可以使用 var os string = runtime.GOOS 来声明 os 的类型，这里使用了自动推断类型
	var os = runtime.GOOS
	if len(os) == 0 {
		t.Error("os is empty")
	}

	// var p = &os - 你可以使用 & 来取得变量存储的地址
	var p = &os
	var q = p
	if *p != os {
		t.Error("wrong assignment")
	}
	if *q != os {
		t.Error("wrong assignment")
	}

	// 修改 p 的指向的值相当于修改 q 指向的值，因为两个相当于同一个
	// 注意 *q != os - 字符串是引用
	*p = "new string"
	// 此时 *p == *q == os
	if strings.Compare(*p, os) != 0 {
		t.Error("modification failed")
	}
}

func TestLocalVariablesOne(t *testing.T) {
	var a, b int
	var c string
	// Go 支持多重赋值
	a, b, c = 5, 7, "abc"
	assert.Equal(t, 5, a)
	assert.Equal(t, 7, b)
	assert.Equal(t, "abc", c)
	// 交换 a 和 b
	a, b = b, a
	assert.Equal(t, 7, a)
	assert.Equal(t, 5, b)
}

func TestLocalVariablesTwo(t *testing.T) {
	// 你可以直接使用 := 赋值而无需声明类型，但这只能用于函数体内，而不可以用于全局变量
	// 使用 := 声明并赋值时，左边的变量必须是首次声明，函数体内的变量一旦声明，则必须被使用
	a, b, c := 5, 7, "abc"
	assert.Equal(t, 5, a)
	assert.Equal(t, 7, b)
	assert.Equal(t, "abc", c)
}

func TestLocalVariablesThree(t *testing.T) {
	// 使用 := 进行并行赋值时，左边只需满足至少一个是第一次声明的即可
	// 比如下面的 a 已经声明过，但还是可以使用 := 操作符
	var a int
	a, b, c := 1, 2, 3
	assert.Equal(t, 1, a)
	assert.Equal(t, 2, b)
	assert.Equal(t, 3, c)
}

func TestVariableInitializedInInit(t *testing.T) {
	assert.Equal(t, 0.7853981633974483, declaredVariable)
}

func TestBlankIdentifier(t *testing.T) {
	// 使用 _ 你可以丢弃掉某些返回值 - 在 Go 中，函数是支持多个返回值，此时你可以使用 _ 来获得自己想要的部分返回值
	// _ 是一个只写变量
	var _, age = "Ray", 24
	assert.Equal(t, 24, age)
}

func TestGetName(t *testing.T) {
	// 使用 _ 选择性丢弃某些返回值
	_, middle, _ := GetName()
	assert.Equal(t, "Young", middle)
}

func ExampleIterateOverAString() {
	IterateOverAString()
	// Output:
	// The index number of H is 0
	// The index number of e is 1
	// The index number of l is 2
	// The index number of l is 3
	// The index number of o is 4
	// The index number of   is 5
	// The index number of W is 6
	// The index number of o is 7
	// The index number of r is 8
	// The index number of l is 9
	// The index number of d is 10
}
