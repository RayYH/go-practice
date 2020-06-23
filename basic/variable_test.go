package basic

import (
	"fmt"
	"runtime"
	"strings"
	"testing"
)

// 全局变量可以在函数内使用
func ExampleShowGlobalVariableOne() {
	fmt.Println("global:", globalString)
	// Output: global: This is a string.
}

func ExampleShowGlobalVariableTwo() {
	fmt.Println(myName, myAge)
	// Output: Ray 24
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

func ExampleShowLocalVariableTwo() {
	var a, b int
	var c string
	a, b, c = 5, 7, "abc"
	fmt.Println("a b c:", a, b, c)
	// Output: a b c: 5 7 abc
}

func ExampleShowLocalVariableThree() {
	// 你可以直接使用 := 赋值而无需声明类型，但这只能用于函数体内，而不可以用于全局变量
	// 使用 := 声明并赋值时，左边的变量必须是首次声明，函数体内的变量一旦声明，则必须被使用
	a, b, c := 5, 7, "abc"
	fmt.Println("a b c:", a, b, c)
	// Output: a b c: 5 7 abc
}

func ExampleShowVariableInitializedInInit() {
	fmt.Println(declaredVariable)
	// Output: 0.7853981633974483
}

func ExampleBlankIdentifier() {
	// 使用 _ 你可以丢弃掉某些返回值 - 在 Go 中，函数是支持多个返回值，此时你可以使用 _ 来获得自己想要的部分返回值
	// _ 是一个只写变量
	var _, age = "Ray", 24
	fmt.Println("age:", age)
	// Output: age: 24
}
