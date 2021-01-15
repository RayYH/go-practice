package oop

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFunctionsWithAPointerArgumentMustTakeAPointer(t *testing.T) {
	v := Vertex{3, 4}
	ScaleFunc(&v, 10)
	assert.Equal(t, v.X, float64(30))
	assert.Equal(t, v.Y, float64(40))

	p := &Vertex{4, 3}
	ScaleFunc(p, 8)
	assert.Equal(t, p.X, float64(32))
	assert.Equal(t, p.Y, float64(24))
}

func TestMethodsWithValueReceiversTakeEitherAValueOrAPointerAsTheReceiver(t *testing.T) {
	v := Vertex{3, 4}
	v.Scale(10)
	assert.Equal(t, v.X, float64(30))
	assert.Equal(t, v.Y, float64(40))

	p := &Vertex{4, 3}
	p.Scale(8)
	assert.Equal(t, p.X, float64(32))
	assert.Equal(t, p.Y, float64(24))
}

// Under the hood, interface values can be thought of as a tuple of a value and a concrete type: (value, type)
// An interface value holds a value of a specific underlying concrete type.
// Calling a method on an interface value executes the method of the same name on its underlying type.
func TestAValueOfInterfaceTypeCanHoldAnyValueThatImplementMethods(t *testing.T) {
	var a Abser
	// A nil interface value holds neither value nor concrete type.
	assert.Equal(t, fmt.Sprintf("(%v, %T)", a, a), "(<nil>, <nil>)")
	f := MyFloat64(-2.0)
	v := Vertex{3.0, 4.0}
	a = f // a MyFloat implements Abser
	assert.Equal(t, a.Abs(), float64(2))
	assert.Equal(t, fmt.Sprintf("(%v, %T)", a, a), "(-2, oop.MyFloat64)")
	// v is a Vertex (not *Vertex) and does NOT implement Abser
	a = &v // a *Vertex implements Abser
	assert.Equal(t, a.Abs(), float64(5))
	assert.Equal(t, fmt.Sprintf("(%v, %T)", a, a), "(&{3 4}, *oop.Vertex)")
}
