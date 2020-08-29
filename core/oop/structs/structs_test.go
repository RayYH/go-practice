package structs

import (
	"fmt"
	"github.com/rayyh/go-practice/core/oop/structs/person"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeclaringStruct(t *testing.T) {
	// 可以通过 var 声明一个结构体
	// var s SimpleStruct 会给 s 分配内存，并零值化内存，这个时候 s 是类型 SimpleStructure
	// 此时 s 被称为 SimpleStruct 的一个实例
	var s SimpleStruct
	s.a = 1
	s.b = 2
	assert.Equal(t, s.a, 1)
	assert.Equal(t, s.b, 2)
	// 结构体是值类型，所以可以通过 new 函数来构造
	// new 返回已分配内存的指针
	// 此时 simpleStruct 是一个指针
	simpleStruct := new(SimpleStruct)
	simpleStruct.a = 1
	simpleStruct.b = 2
	assert.Equal(t, simpleStruct.a, 1)
	assert.Equal(t, simpleStruct.b, 2)
	assert.Equal(t, fmt.Sprintf("%v", simpleStruct), "&{1 2}")
}

func TestStructLiteral(t *testing.T) {
	// 此时 s 类型是 *SimpleStruct，&{} 的底层仍会调用 new()
	s := &SimpleStruct{1, 2}
	assert.Equal(t, s.a, 1)
	assert.Equal(t, s.b, 2)

	var s2 SimpleStruct
	s2 = SimpleStruct{1, 2}
	assert.Equal(t, s2.a, 1)
	assert.Equal(t, s2.b, 2)

	interval1 := Interval{0, 3}
	assert.Equal(t, interval1.start, 0)
	assert.Equal(t, interval1.end, 3)
	interval2 := Interval{end: 5, start: 1}
	assert.Equal(t, interval2.end, 5)
	assert.Equal(t, interval2.start, 1)
	interval3 := Interval{end: 5}
	assert.Equal(t, interval3.start, 0)
	assert.Equal(t, interval3.end, 5)
}

func TestPersonStruct(t *testing.T) {
	// 1-struct as a value type
	var person1 Person
	person1.firstName = "Ray"
	person1.lastName = "Hong"
	upPerson(&person1)
	assert.Equal(t, person1.firstName, "RAY")
	assert.Equal(t, person1.lastName, "HONG")

	// 2—struct as a pointer
	person2 := new(Person)
	person2.firstName = "Ray"
	person2.lastName = "Hong"
	upPerson(person2)
	assert.Equal(t, person2.firstName, "RAY")
	assert.Equal(t, person2.lastName, "HONG")

	// 3—struct as a literal
	person3 := &Person{firstName: "Ray", lastName: "Hong"}
	upPerson(person3)
	assert.Equal(t, person3.firstName, "RAY")
	assert.Equal(t, person3.lastName, "HONG")
}

func TestStructFactory(t *testing.T) {
	f := NewFile(1, "filename")
	assert.Equal(t, f.fd, 1)
	assert.Equal(t, f.name, "filename")
}

func ExampleRefTag() {
	tt := TagType{true, "Ray Hong", 1}
	for i := 0; i < 3; i++ {
		RefTag(tt, i)
	}
	// Output:
	// An important answer
	// The name of the thing
	// How much there are
}

func TestAnonymousField(t *testing.T) {
	outer := new(outerS)
	outer.b = 6
	outer.c = 7.5
	outer.int = 60
	// 外层结构体通过 outer.in1 直接进入内层结构体的字段
	outer.in1 = 5
	outer.in2 = 10

	assert.Equal(t, outer.b, 6)
	assert.Equal(t, outer.c, 7.5)
	assert.Equal(t, outer.int, 60)
	assert.Equal(t, outer.in1, 5)
	assert.Equal(t, outer.in2, 10)

	// using literal
	outer2 := &outerS{6, 7.5, 60, innerS{5, 10}}
	assert.Equal(t, fmt.Sprintf("%v", outer2), "&{6 7.5 60 {5 10}}")
}

func TestEmbeddedStruct(t *testing.T) {
	b := &B{A{1, 2}, 3.0, 4.0}
	assert.Equal(t, b.ax, 1)
	assert.Equal(t, b.ay, 2)
	assert.Equal(t, b.bx, 3.0)
	assert.Equal(t, b.by, 4.0)
	assert.Equal(t, b.A, A{1, 2})
}

func TestNameOverridden(t *testing.T) {
	is := &innerStruct{1, 2}
	os := &outerStruct{is: *is, c: 3, a: 100}
	assert.Equal(t, os.a, 100)
}

func TestTwoIntegers(t *testing.T) {
	two := new(TwoIntegers)
	two.a = 12
	two.b = 10
	assert.Equal(t, two.AddThem(), 22)
	assert.Equal(t, two.AddToParam(30), 52)
}

func TestIntVectorSum(t *testing.T) {
	assert.Equal(t, IntVector{1, 2, 3}.Sum(), 6)
}

func TestIntContainer(t *testing.T) {
	var t1 IntContainer // 值
	t1.change()
	assert.Equal(t, t1.write(), "{1}")
	t2 := new(IntContainer) // 指针
	t2.change()
	assert.Equal(t, t2.write(), "{1}")
}

func TestPersonInPersonPackage(t *testing.T) {
	p := new(person.Person)
	p.SetFirstName("Rayyh")
	assert.Equal(t, "Rayyh", p.FirstName())
}

func TestInheritance(t *testing.T) {
	n := &NamedPoint{Point{3, 4}, "Pythagoras"}
	assert.Equal(t, n.Abs(), float64(5))
}

func ExampleCustomerWithRef_Log() {
	c := new(CustomerWithRef)
	c.Name = "Ray Hong"
	c.log = new(Log)
	c.log.msg = "1 - Yes we can!"
	c = &CustomerWithRef{"Ray Hong", &Log{"1 - Yes we can!"}}
	c.Log().Add("2 - After me the world will be a better place!")
	fmt.Println(c.Log())
	// Output:
	// 1 - Yes we can!
	// 2 - After me the world will be a better place!
}

func ExampleCustomer_String() {
	c := &Customer{"Ray Hong", Log{"1 - Yes we can!"}}
	c.Add("2 - After me the world will be a better place!")
	fmt.Println(c)
	// Output:
	// Ray Hong
	// Log:{1 - Yes we can!
	// 2 - After me the world will be a better place!}
}

func TestMultiInheritance(t *testing.T) {
	cp := new(CameraPhone)
	assert.Equal(t, cp.Call(), "Ring Ring")
	assert.Equal(t, cp.TakeAPicture(), "Click")
}
