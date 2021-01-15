package oop

import (
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
