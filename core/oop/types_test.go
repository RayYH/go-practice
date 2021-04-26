package oop

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Integer int

func (a Integer) LessThan(b Integer) bool {
	return a < b
}

// if you want to modify object `a`
func (a *Integer) Add(b Integer) {
	*a += b
}

func TestIntegerType(t *testing.T) {
	var a Integer = 2
	var b Integer = 3
	assert.Equal(t, true, a.LessThan(b))
	a.Add(b)
	assert.Equal(t, Integer(5), a)
}

func TestArrayIsOfValueType(t *testing.T) {
	var a = [3]int{1, 2, 3}
	var b = a
	b[0] = -1
	assert.Equal(t, a, [3]int{1, 2, 3})
	assert.Equal(t, b, [3]int{-1, 2, 3})

	var c = [3]int{1, 2, 3}
	var d = &c
	d[0] = -1
	assert.Equal(t, c, [3]int{-1, 2, 3})
	assert.Equal(t, *d, [3]int{-1, 2, 3})
}
