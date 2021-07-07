package advanced_types

import (
	"fmt"
	"reflect"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

func TestStructsDefinitions(t *testing.T) {
	t.Run("basic struct declaration", func(t *testing.T) {
		type IntContainer struct{ thing int }
		t1 := IntContainer{thing: 1}

		assert.Equal(t, 1, t1.thing)
	})

	t.Run("anonymous field", func(t *testing.T) {
		type BooleanContainer struct {
			bool // anonymous field
		}

		b := BooleanContainer{}
		b.bool = true
		assert.True(t, b.bool)
	})
}

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

	t.Run("access inner anonymous fields directly", func(t *testing.T) {
		s := new(outerStruct)
		s.intValue = 1
		assert.Equal(t, s.intValue, 1)
	})

	t.Run("access inner fields via chaining invocation", func(t *testing.T) {
		s := new(outerStruct)
		s.bs.boolValue = false
		assert.Equal(t, s.bs.boolValue, false)
	})

	t.Run("name overridden", func(t *testing.T) {
		f := floatStruct{floatValue: 1.1}
		o := &outerStruct{
			floatStruct: f,
			floatValue:  2.2,
		}
		assert.Equal(t, o.floatValue, 2.2)
		assert.Equal(t, f.floatValue, 1.1)
	})
}

func TestStructsDeclaration(t *testing.T) {
	t.Run("var", func(t *testing.T) {
		var i1 Interval
		assert.Equal(t, i1.start, 0)
		assert.Equal(t, i1.end, 0)

		i1.start, i1.end = 1, 100
		assert.Equal(t, i1.start, 1)
		assert.Equal(t, i1.end, 100)
		assert.Equal(t, "{1 100}", fmt.Sprintf("%v", i1))
	})

	t.Run("new", func(t *testing.T) {
		i2 := new(Interval)
		i2.start, i2.end = 1, 100
		assert.Equal(t, i2.start, 1)
		assert.Equal(t, i2.end, 100)
		assert.Equal(t, "&{1 100}", fmt.Sprintf("%v", i2))
	})
}

func TestStructsLiterals(t *testing.T) {
	var i1 Interval
	i1 = Interval{0, 1} // {0 1}
	assert.Equal(t, i1.start, 0)
	assert.Equal(t, i1.end, 1)

	i2 := Interval{0, 3}
	assert.Equal(t, i2.start, 0)
	assert.Equal(t, i2.end, 3)

	i3 := &Interval{0, 2} // &{0 2}
	assert.Equal(t, i3.start, 0)
	assert.Equal(t, i3.end, 2)

	i4 := Interval{end: 4, start: 1}
	assert.Equal(t, i4.start, 1)
	assert.Equal(t, i4.end, 4)

	i5 := Interval{end: 5}
	assert.Equal(t, i5.start, 0)
	assert.Equal(t, i5.end, 5)
}

func TestStructFeatures(t *testing.T) {
	t.Run("both value type and pointer type work on structs", func(t *testing.T) {
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

	t.Run("struct with tags", func(t *testing.T) {
		type MixinContainer struct {
			BoolValue   bool   `description:"a bool value"`
			StringValue string `description:"a string value"`
			IntValue    int    `description:"an int value"`
		}

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

		field, found := reflect.TypeOf(mixinContainer).FieldByName("StringValue")
		if found {
			assert.Equal(t, "a string value", fmt.Sprintf("%v", field.Tag.Get("description")))
		}
	})

	//
	t.Run("struct recursion", func(t *testing.T) {
		type Node struct {
			prev *Node
			data int
			next *Node
		}

		head := new(Node)
		head.prev = nil
		head.data = 0
		head.next = new(Node)
		// more nodes...
	})

	t.Run("struct size", func(t *testing.T) {
		type T struct {
			t float64
		}
		assert.Equal(t, uintptr(0x8), unsafe.Sizeof(T{}))
	})
}

func TestComparing2OOP(t *testing.T) {
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

	t.Run("inheritance", func(t *testing.T) {
		t.Run("CameraPhone", func(t *testing.T) {
			cp := new(CameraPhone)
			assert.Equal(t, cp.Call(), "Ring Ring")
			assert.Equal(t, cp.TakeAPicture(), "Click")
		})

		t.Run("NamedPoint", func(t *testing.T) {
			n := &NamedPoint{Point{3, 4}, "Pythagoras"}
			assert.Equal(t, n.Abs(), float64(5))
		})
	})

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

	t.Run("factory methods", func(t *testing.T) {
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

func TestLog(t *testing.T) {
	t.Run("Customer", func(t *testing.T) {
		c := new(Customer)
		c.Name = "Ray Hong"
		c.Log = Log{"first message"}
		c.Add("second message")
		assert.Equal(t, "Ray Hong\nLog:{first message\nsecond message}\n", fmt.Sprint(c))
	})

	t.Run("CustomerHoldsRef", func(t *testing.T) {
		c1 := new(CustomerHoldsRef)
		c1.Name = "Ray Hong"
		c1.log = new(Log)
		c1.log.msg = "first message"
		assert.Equal(t, fmt.Sprint(c1.Log()), "first message")

		c2 := CustomerHoldsRef{
			Name: "Ray Hong",
			log:  &Log{"first message"},
		}
		c2.Log().Add("second message")
		assert.Equal(t, fmt.Sprint(c2.Log()), "first message\nsecond message")
	})
}
