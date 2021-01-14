package oop

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInheritance(t *testing.T) {
	n := &NamedPoint{Point{3, 4}, "Pythagoras"}
	// anonymous struct member's method can be invoked directly
	assert.Equal(t, n.Abs(), float64(5))
}

func TestMultipleInheritance(t *testing.T) {
	cp := new(CameraPhone)
	assert.Equal(t, cp.Call(), "Ring Ring")
	assert.Equal(t, cp.TakeAPicture(), "Click")
}
