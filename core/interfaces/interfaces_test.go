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
	// 一个接口类型的变量 varI 中可以包含任何类型的值，必须有一种方式来检测它的动态类型，即运行时在变量中存储的值的实际类型
	// 通常我们可以使用 类型断言 来测试在某个时刻 varI 是否包含类型 T 的值：v := varI.(T)
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
			// 可以用 type-switch 进行运行时类型分析，但是在 type-switch 不允许有 fallthrough
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
		// 这里必须用 *Square 而不是 Square，因为 Area() 方法的接收者是 *Square 而不是 Square，这与方法调用时的智能指针有所不同
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

	// 测试一个值是否实现了某个接口
	t.Run("check interface", func(t *testing.T) {
		var s interface{}
		s = &Rectangle{length: 20, width: 20}

		// 检查 s 是否实现了 Shaper 接口
		if sv, ok := s.(Shaper); ok {
			assert.Equal(t, float64(400), sv.Area())
		}
	})
}

func TestMethodSets(t *testing.T) {
	// 在接口上调用方法时，必须有和方法定义时相同的接收者类型或者是可以从具体类型 P 直接可以辨识的
	// -> 指针方法可以通过指针调用
	// -> 值方法可以通过值调用
	// -> 接收者是值的方法可以通过指针调用，因为指针会首先被解引用
	// -> 接收者是指针的方法不可以通过值调用，因为存储在接口中的值没有地址

	// Go 语言规范定义了接口方法集的调用规则
	// -> 类型 T 的可调用方法集包含接受者为 *T 或 T 的所有方法集
	// -> 类型 *T 的可调用方法集包含接受者为 *T 的所有方法
	// -> 类型 *T 的可调用方法集不包含接受者为 T 的方法
	t.Run("using value type", func(t *testing.T) {
		var lst List
		assert.False(t, LongEnough(lst))
	})

	t.Run("using pointer type", func(t *testing.T) {
		pl := new(List)
		CountInto(pl, 1, 10)
		assert.True(t, LongEnough(pl))
	})
}
