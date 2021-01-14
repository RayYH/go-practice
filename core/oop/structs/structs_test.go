package structs

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

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
