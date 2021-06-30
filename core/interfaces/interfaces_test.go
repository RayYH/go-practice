package interfaces

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEmptyInterfaceHoldsValueOfAnyType(t *testing.T) {
	var i interface{}
	assert.Equal(t, fmt.Sprintf("(%v, %T)", i, i), "(<nil>, <nil>)")
	i = 42
	assert.Equal(t, fmt.Sprintf("(%v, %T)", i, i), "(42, int)")
	i = "hello"
	assert.Equal(t, fmt.Sprintf("(%v, %T)", i, i), "(hello, string)")
}

// i.(type) 可以检查变量是否为指定类型
func TestTypeAssertions(t *testing.T) {
	t.Run("type determination", func(t *testing.T) {
		var i interface{} = "hello"
		s := i.(string)
		assert.Equal(t, s, "hello")
		s, ok := i.(string)
		assert.Equal(t, s, "hello")
		assert.True(t, ok)
		_, ok = i.(float64) // hello cannot be converted to float64 type
		assert.False(t, ok)
	})

	t.Run("type switch", func(t *testing.T) {
		classifier := func(items ...interface{}) (res string) {
			// A type switch is a construct that permits several type assertions in series.
			for i, x := range items {
				switch x.(type) {
				case bool:
					res += fmt.Sprintf("[#%d: bool]", i)
				case float64:
					res += fmt.Sprintf("[#%d: float64]", i)
				case int, int64:
					res += fmt.Sprintf("[#%d: int]", i)
				case nil:
					res += fmt.Sprintf("[#%d: nil]", i)
				case string:
					res += fmt.Sprintf("[#%d: string]", i)
				default:
					res += fmt.Sprintf("[#%d: unknown]", i)
				}
			}
			return
		}

		assert.Equal(t, "[#0: int][#1: float64][#2: string][#3: unknown][#4: nil][#5: bool]", classifier(13, -14.3, "BELGIUM", complex(1, 2), nil, false))
	})
}

func TestAbser(t *testing.T) {
	var a Abser
	assert.Equal(t, "(<nil>, <nil>)", fmt.Sprintf("(%v, %T)", a, a))

	f := MyFloat64(-2.0)
	v := Vertex{3.0, 4.0}

	a = f // a MyFloat implements Abser
	assert.Equal(t, float64(2), a.Abs())
	assert.Equal(t, "(-2, interfaces.MyFloat64)", fmt.Sprintf("(%v, %T)", a, a))

	// 智能指针
	a = &v // a *Vertex (v Vertex) implements Abser
	assert.Equal(t, float64(5), a.Abs())
	assert.Equal(t, float64(5), v.Abs())
	assert.Equal(t, "(&{3 4}, *interfaces.Vertex)", fmt.Sprintf("(%v, %T)", a, a))
	assert.Equal(t, "({3 4}, interfaces.Vertex)", fmt.Sprintf("(%v, %T)", v, v))
}

func TestShaper(t *testing.T) {
	t.Run("area method", func(t *testing.T) {
		s := &Square{5.0}
		r := &Rectangle{length: 5.0, width: 5.0}

		// Shaper 切片可以容纳 Square 和 Rectangle 类型，因为这两个结构体都实现了 Shaper 接口包含的方法集
		shapes := []Shaper{s, r}

		for _, shape := range shapes {
			assert.Equal(t, 25.0, shape.Area())
		}
	})

	t.Run("check type", func(t *testing.T) {
		var s Shaper
		s = new(Rectangle)
		_, ok := s.(*Square)
		assert.False(t, ok)
		_, ok = s.(*Rectangle)
		assert.True(t, ok)
	})

	t.Run("get type", func(t *testing.T) {
		var s1, s2 Shaper
		s1 = &Rectangle{length: 1, width: 2}
		s2 = &Square{side: 1}
		assert.Equal(t, "Rec", GetType(s1))
		assert.Equal(t, "Squ", GetType(s2))
	})
}
