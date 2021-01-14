package oop

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Integer int

func (a Integer) LessThan(b Integer) bool {
	return a < b
}

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
