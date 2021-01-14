package oop

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStructsDeclaration(t *testing.T) {
	// declare a struct through var keyword
	// allocate memory and initialize members with zero values!
	var i1 Interval
	assert.Equal(t, i1.start, 0)
	assert.Equal(t, i1.end, 0)
	// modify values
	i1.start = 1
	i1.end = 100
	assert.Equal(t, i1.start, 1)
	assert.Equal(t, i1.end, 100)
	assert.Equal(t, "{1 100}", fmt.Sprintf("%v", i1))

	// since struct is a value type, we can use new keyword
	i2 := new(Interval)
	i2.start = 1
	i2.end = 100
	assert.Equal(t, i2.start, 1)
	assert.Equal(t, i2.end, 100)
	assert.Equal(t, "&{1 100}", fmt.Sprintf("%v", i2))
}

func TestStructsLiterals(t *testing.T) {
	var i1 Interval
	i1 = Interval{0, 1} // {0 1}
	assert.Equal(t, i1.start, 0)
	assert.Equal(t, i1.end, 1)

	i2 := &Interval{0, 2} // &{0 2}
	assert.Equal(t, i2.start, 0)
	assert.Equal(t, i2.end, 2)

	i3 := Interval{0, 3}
	assert.Equal(t, i3.start, 0)
	assert.Equal(t, i3.end, 3)

	i4 := Interval{end: 4, start: 1}
	assert.Equal(t, i4.end, 4)
	assert.Equal(t, i4.start, 1)

	i5 := Interval{end: 5}
	assert.Equal(t, i5.start, 0)
	assert.Equal(t, i5.end, 5)
}

func TestIntervalMethods(t *testing.T) {
	i := &Interval{0, 10}
	assert.Equal(t, i.duration(), 10)
	assert.Equal(t, i.durationInMillSeconds(), 10*1000)
}
