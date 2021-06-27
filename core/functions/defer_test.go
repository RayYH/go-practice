package functions

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// 关键字 defer 的用法类似于面向对象编程语言 Java 和 C# 的 finally 语句块，它一般用于释放某些已分配的资源
func ExampleGreet() {
	defer Greet()
	fmt.Println("DEFER")
	// Output:
	// DEFER
	// Hello
}

// Go 中 defer 是以入栈的方式保存的，因此调用顺序符合先进后出的规则
func ExampleDeferOrders() {
	defer DeferOrders()
	// Output:
	// 3
	// 2
	// 1
}

func TestNamedReturnValuesWithDefer(t *testing.T) {
	getValue := func() (ret int) {
		defer func() {
			ret++
		}()
		// 这里的 return 1 相当于 ret = 1 和 return ret 的组合，因此后面 defer 可以继续更改 ret 的值
		return 1
	}

	assert.Equal(t, getValue(), 2)

	// 如果我们的返回值没有命名，则 defer 操作的只是一个函数内部的变量 ret
	getValueWithoutNamedReturn := func() int {
		ret := 0
		defer func() {
			ret++
		}()
		return 1
	}
	assert.Equal(t, getValueWithoutNamedReturn(), 1)
}
