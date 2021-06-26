package basic_types

import (
	"fmt"
	"math"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Python 中的类型转换与其他语言类似，使用 `T(v)` 即可将值 `v` 强转为类型 `T`
func TestTypeConversions(t *testing.T) {
	var n int16 = 34
	var m int32
	// 这时 m 已经被转换为 int32 了，与 int16 的 n 不同
	m = int32(n)
	assert.NotEqual(t, reflect.TypeOf(n), reflect.TypeOf(m))
	// int32(n) and m are of the same type and same value
	assert.Equal(t, reflect.TypeOf(int32(n)), reflect.TypeOf(m))

	var x, y = 3, 4
	var f = math.Sqrt(float64(x*x + y*y))
	var z = uint(f)
	var r uint = 5
	assert.Equal(t, r, z)
}

func TestTypeInterface(t *testing.T) {
	i := 42
	j := 3.14
	k := 0.867 + 0.5i
	// %T 打印出值的类型
	assert.Equal(t, "int float64 complex128", fmt.Sprintf("%T %T %T", i, j, k))
}

func TestTypeAliases(t *testing.T) {
	// alias types
	type (
		A1 = string
		A2 = A1
	)

	// custom types
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

	assert.Equal(t, "", a1)
	assert.Equal(t, "", a2)
	// 注意下面不能用 ""，而是必须强转为指定的类型
	assert.Equal(t, B1(""), b1)
	assert.Equal(t, B2(""), b2)
	assert.Nil(t, b3)
	assert.Nil(t, b4)
}
