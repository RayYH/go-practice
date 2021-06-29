package advanced_types

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStructsDefinitions(t *testing.T) {
	// 使用 type [struct_name] struct { } 可以定义一个结构体
	t.Run("basic struct declaration", func(t *testing.T) {
		type IntContainer struct{ thing int }
		t1 := IntContainer{thing: 1}

		assert.Equal(t, 1, t1.thing)
	})

	// Go 中的结构体可以不用声明 field 的名称，只声明 field 的类型，此时我们称这个 field 为匿名 field
	t.Run("anonymous field", func(t *testing.T) {
		type BooleanContainer struct {
			bool // anonymous field
		}

		b := BooleanContainer{}
		b.bool = true
		assert.True(t, b.bool)
	})
}

// Go 中结构体成员的类型可以是另一个结构体，这种我们称之为内嵌结构体
func TestEmbeddingStructs(t *testing.T) {
	type intStruct struct{ intValue int }
	type boolStruct struct{ boolValue bool }
	type floatStruct struct{ floatValue float64 }
	type outerStruct struct {
		intStruct              // anonymous struct field intStruct
		floatStruct            // anonymous struct field floatStruct
		bs          boolStruct // named struct field
		floatValue  float64    // override inner struct
		stringValue string     // self field
	}

	// 对于匿名的结构体属性，无需使用 outerStruct.innerStruct.Field 这种语法，直接使用 outerStruct.Field 即可
	t.Run("access inner anonymous fields directly", func(t *testing.T) {
		s := new(outerStruct)
		s.intValue = 1
		assert.Equal(t, s.intValue, 1)
	})

	// 链式调用获取内部结构体中成员的值
	t.Run("access inner fields via chaining invocation", func(t *testing.T) {
		s := new(outerStruct)
		s.bs.boolValue = false
		assert.Equal(t, s.bs.boolValue, false)
	})

	// 外部结构体中的 field 可以覆盖 (隐藏) 内部结构体中的同名 field
	t.Run("name overridden", func(t *testing.T) {
		f := floatStruct{floatValue: 1.1}
		o := &outerStruct{
			floatStruct: f,
			floatValue:  2.2,
		}
		assert.Equal(t, o.floatValue, 2.2)
		assert.Equal(t, f.floatValue, 1.1)
		// 这里无法通过链式调用来获取 floatStruct 中的 floatValue 的值，因为 floatStruct 是一个匿名成员
	})
}

func TestStructsDeclaration(t *testing.T) {
	t.Run("var", func(t *testing.T) {
		// 可以通过 `var` 关键字声明一个结构体，一旦结构体被声明，其所有的 field 都会被初始化为零值
		// 访问结构体实例中的成员需要使用 `.` 语法
		var i1 Interval
		assert.Equal(t, i1.start, 0)
		assert.Equal(t, i1.end, 0)

		// 声明了实例之后，我们可以修改实例中 field 的值
		i1.start, i1.end = 1, 100
		assert.Equal(t, i1.start, 1)
		assert.Equal(t, i1.end, 100)
		assert.Equal(t, "{1 100}", fmt.Sprintf("%v", i1))
	})

	t.Run("new", func(t *testing.T) {
		// 我们还可以使用 new() 方法来初始化一个结构体，这时返回的是一个指针类型
		i2 := new(Interval)
		i2.start, i2.end = 1, 100
		assert.Equal(t, i2.start, 1)
		assert.Equal(t, i2.end, 100)
		assert.Equal(t, "&{1 100}", fmt.Sprintf("%v", i2))
	})
}

func TestStructsLiterals(t *testing.T) {
	// 常规写法，先声明、再赋值
	var i1 Interval
	i1 = Interval{0, 1} // {0 1}
	assert.Equal(t, i1.start, 0)
	assert.Equal(t, i1.end, 1)

	// `:=` 运算符直接初始化一个结构体
	i2 := Interval{0, 3}
	assert.Equal(t, i2.start, 0)
	assert.Equal(t, i2.end, 3)

	// `:=` 运算符配合 & 运算符，此时 i3 是一个指针
	i3 := &Interval{0, 2} // &{0 2}
	assert.Equal(t, i3.start, 0)
	assert.Equal(t, i3.end, 2)

	// 在字面量中指名 field 的名称
	i4 := Interval{end: 4, start: 1}
	assert.Equal(t, i4.end, 4)
	assert.Equal(t, i4.start, 1)

	// 只初始化部分成员，其他成员为对应类型的零值
	i5 := Interval{end: 5}
	assert.Equal(t, i5.start, 0)
	assert.Equal(t, i5.end, 5)
}

func TestStructFeatures(t *testing.T) {
	// 无论是值类型，还是指针类型，Go 都可以使用 `.` 语法来调用结构体可用的方法，这种特性被称为智能指针
	// 简而言之：指针方法和值方法都可以在指针或非指针上被调用
	t.Run("both value type and pointer type work on structs (smart pointers)", func(t *testing.T) {
		// value type
		var t1 IntContainer
		t1.change()
		assert.Equal(t, "{1}", t1.toString())
		assert.Equal(t, 2, t1.modify())
		assert.Equal(t, "{1}", t1.toString())

		// pointer type
		t2 := new(IntContainer)
		t2.change()
		assert.Equal(t, "{1}", t2.toString())
		assert.Equal(t, 2, t2.modify())
		assert.Equal(t, "{1}", t2.toString())
	})

	// Go 中的结构体可以为每个 field 添加 tag
	t.Run("struct with tags", func(t *testing.T) {
		type MixinContainer struct {
			BoolValue   bool   `description:"a bool value"`
			StringValue string `description:"a string value"`
			IntValue    int    `description:"an int value"`
		}

		// 根据指定的 index 获取对应 field 的标签中的 description 值
		inspectField := func(mc MixinContainer, index int) string {
			r := reflect.TypeOf(mc)
			field := r.Field(index)
			return fmt.Sprintf("%v", field.Tag.Get("description"))
		}

		mixinContainer := MixinContainer{
			BoolValue:   true,
			StringValue: "Ray Hong",
			IntValue:    1,
		}
		assert.Equal(t, "a bool value", inspectField(mixinContainer, 0))
		assert.Equal(t, "a string value", inspectField(mixinContainer, 1))
		assert.Equal(t, "an int value", inspectField(mixinContainer, 2))

		// 一种更为常用的反射用法是根据结构体中 field 的 Name 来获取对应 Tag 中的值
		field, found := reflect.TypeOf(mixinContainer).FieldByName("StringValue")
		if found {
			assert.Equal(t, "a string value", fmt.Sprintf("%v", field.Tag.Get("description")))
		}
	})
}

func TestComparing2OOP(t *testing.T) {
	// 如果把 Go 中的结构体类比为其他语言中的类，由于 Go 中没有构造方法，因此我们通常使用 NewXXX 来初始化指定结构体，我们称这类方法为工厂方法
	t.Run("factory methods", func(t *testing.T) {
		type File struct {
			fd   int
			name string
		}

		NewFile := func(fd int, name string) *File {
			if fd < 0 {
				return nil
			}

			return &File{fd, name}
		}

		f := NewFile(1, "filename")
		assert.Equal(t, f.fd, 1)
		assert.Equal(t, f.name, "filename")
	})

	// Go 不支持 extends 关键字来实现继承，Go 是通过组合的方式来达到继承的效果的，有点类似于 JS/TS 中的 duck typing
	t.Run("inheritance", func(t *testing.T) {
		t.Run("CameraPhone", func(t *testing.T) {
			// CameraPhone 是由 Camera 结构体和 Phone 结构体组合而成
			cp := new(CameraPhone)
			assert.Equal(t, cp.Call(), "Ring Ring")
			assert.Equal(t, cp.TakeAPicture(), "Click")
		})

		t.Run("NamedPoint", func(t *testing.T) {
			// Abs() 方法定义在 Point 结构体中，Point 结构体是 NamedPoint 的匿名成员，因此 n 可以直接调用 Abs() 方法
			n := &NamedPoint{Point{3, 4}, "Pythagoras"}
			assert.Equal(t, n.Abs(), float64(5))
		})
	})

	// Go 中哪些类型拥有哪些方法是非常灵活的，只需要在方法声明时指名方法的调用者类型即可
	t.Run("methods", func(t *testing.T) {
		t.Run("IntVector", func(t *testing.T) {
			assert.Equal(t, IntVector{1, 2, 3}.Sum(), 6)
		})

		t.Run("Integer", func(t *testing.T) {
			var a Integer = 2
			var b Integer = 3
			assert.Equal(t, true, a.LessThan(b))
			a.Add(b)
			assert.Equal(t, Integer(5), a)
		})
	})
}

func TestFunctionsComparing2Methods(t *testing.T) {
	// 对于函数而言，指针类型的参数必须接受指针类型的变量
	t.Run("functions", func(t *testing.T) {
		v := Vertex{3, 4}
		ScaleFunc(&v, 10)
		assert.Equal(t, v.X, float64(30))
		assert.Equal(t, v.Y, float64(40))

		p := &Vertex{4, 3}
		ScaleFunc(p, 8)
		assert.Equal(t, p.X, float64(32))
		assert.Equal(t, p.Y, float64(24))
	})

	// 对于方法而言，无论调用者是指针类型还是值类型，代码都能正常运行，这种机制称为智能指针
	t.Run("methods", func(t *testing.T) {
		v := Vertex{3, 4}
		v.Scale(10)
		assert.Equal(t, v.X, float64(30))
		assert.Equal(t, v.Y, float64(40))

		p := &Vertex{4, 3}
		p.Scale(8)
		assert.Equal(t, p.X, float64(32))
		assert.Equal(t, p.Y, float64(24))
	})
}

// Person 结构体示例
func TestPerson(t *testing.T) {
	t.Run("value type", func(t *testing.T) {
		var person1 Person
		person1.firstName, person1.lastName = "Ray", "Hong"
		upperPerson(&person1)
		assert.Equal(t, person1.firstName, "RAY")
		assert.Equal(t, person1.lastName, "HONG")
	})

	t.Run("pointer type", func(t *testing.T) {
		person2 := new(Person) // new allocates the memory required by a type
		person2.firstName, person2.lastName = "Ray", "Hong"
		upperPerson(person2)
		assert.Equal(t, person2.firstName, "RAY")
		assert.Equal(t, person2.lastName, "HONG")
	})

	t.Run("literal", func(t *testing.T) {
		person3 := &Person{firstName: "Ray", lastName: "Hong"}
		upperPerson(person3)
		assert.Equal(t, person3.firstName, "RAY")
		assert.Equal(t, person3.lastName, "HONG")
	})

	t.Run("constructor", func(t *testing.T) {
		person4 := NewPerson("Ray", "Hong")
		upperPerson(person4)
		assert.Equal(t, person4.firstName, "RAY")
		assert.Equal(t, person4.lastName, "HONG")
	})

	t.Run("string format", func(t *testing.T) {
		person := &Person{firstName: "Ray", lastName: "Hong"}
		assert.Equal(t, fmt.Sprint(person), "Ray Hong")
	})
}
