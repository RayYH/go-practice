package interfaces

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

// 在 Go 的底层，接口可以被认为一个是一个包含了值和类型的元组：(value, type)
// 接口值保存了一个具体底层类型的具体值
// 接口值调用方法时会执行其底层类型的同名方法
func TestAValueOfInterfaceTypeCanHoldAnyValueThatImplementMethods(t *testing.T) {
	var a Abser
	// A nil interface value holds neither value nor concrete type.
	assert.Equal(t, fmt.Sprintf("(%v, %T)", a, a), "(<nil>, <nil>)")
	f := MyFloat64(-2.0)
	v := Vertex{3.0, 4.0}

	a = f // a MyFloat implements Abser
	assert.Equal(t, a.Abs(), float64(2))
	assert.Equal(t, fmt.Sprintf("(%v, %T)", a, a), "(-2, interfaces.MyFloat64)")

	// v is a Vertex (not *Vertex) and does NOT implement Abser
	a = &v // a *Vertex implements Abser
	assert.Equal(t, a.Abs(), float64(5))
	assert.Equal(t, fmt.Sprintf("(%v, %T)", a, a), "(&{3 4}, *interfaces.Vertex)")
}
