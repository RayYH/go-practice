package types

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLessThan(t *testing.T) {
	var a Integer = 2
	var b Integer = 3
	assert.Equal(t, true, a.LessThan(b))
	assert.Equal(t, true, LessThan(a, b))
}

func TestIntegerAdd(t *testing.T) {
	var a Integer = 1
	var b Integer = 2
	var c Integer = 3
	a.Add(b)
	assert.Equal(t, c, a)
}
