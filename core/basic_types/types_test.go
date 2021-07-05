package basic_types

import (
	"fmt"
	"math"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Go 中的类型转换与其他语言类似，使用 `T(v)` 即可将值 `v` 强转为类型 `T`
// (Go 是一门静态类型语言，不支持隐式类型转换)
func TestTypeConversions(t *testing.T) {
	var n int16 = 34
	var m int32

	// m 为 int32 类型，但是值与 n 相同
	m = int32(n)
	assert.NotEqual(t, reflect.TypeOf(n), reflect.TypeOf(m)) // 两者类型不同
	assert.Equal(t, fmt.Sprint(m), fmt.Sprint(n))            // 若按字符串输出则都会输出 32

	var x, y = 3, 4
	var f = math.Sqrt(float64(x*x + y*y)) // sqrt 会返回 float64 类型
	var z = uint(f)
	var r uint = 5
	assert.Equal(t, r, z)

	// 当从一个取值范围较大的转换到取值范围较小的类型时，会发生精度丢失的情况
	t.Run("lost precision", func(t *testing.T) {
		f64 := 3.1415926535897932384626433
		f32 := float32(f64)
		assert.Equal(t, fmt.Sprintf("%v %v", f32, f64), "3.1415927 3.141592653589793")
	})
}

// 使用 type 关键字可以定义你自己的类型 (结构体、接口)，也可以定义一个已经存在的类型的别名
// 这里并不是真正意义上的别名，因为使用这种方法定义之后的类型可以拥有更多的特性，且在类型转换时必须显式转换
func TestTypeAliases(t *testing.T) {
	// 使用 = 表示别名，A1，A2，string 本质上是同一个类型
	type (
		A1 = string
		A2 = A1
	)

	// 不使用 = 表示定义一个新类型，新类型不会有原类型所附带的方法
	type (
		B1 string
		B2 B1
		B3 []B1
		B4 B3
	)

	var a1 A1
	var a2 A2
	var b1 B1
	var b2 B2
	var b3 B3
	var b4 B4

	// 别名的话可以和原类型进行比较
	assert.Equal(t, "", a1)
	assert.Equal(t, "", a2)

	// 新类型的话必须转换为指定的类型才可以继续进行比较
	assert.Equal(t, B1(""), b1)
	assert.Equal(t, B2(""), b2)
	assert.Nil(t, b3)
	assert.Nil(t, b4)
}
