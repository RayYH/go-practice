package oop

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func ExampleClassifierCaller() {
	ClassifierCaller()
	// Output:
	// Param #0 is a int
	// Param #1 is a float64
	// Param #2 is a string
	// Param #3 is unknown
	// Param #4 is a nil
	// Param #5 is a bool
}

func TestEmptyInterfaceHoldsValueOfAnyType(t *testing.T) {
	var i interface{}
	assert.Equal(t, fmt.Sprintf("(%v, %T)", i, i), "(<nil>, <nil>)")
	i = 42
	assert.Equal(t, fmt.Sprintf("(%v, %T)", i, i), "(42, int)")
	i = "hello"
	assert.Equal(t, fmt.Sprintf("(%v, %T)", i, i), "(hello, string)")
}

func TestTypeAssertions(t *testing.T) {
	var i interface{} = "hello"
	s := i.(string)
	assert.Equal(t, s, "hello")
	s, ok := i.(string)
	assert.Equal(t, s, "hello")
	assert.True(t, ok)
	_, ok = i.(float64) // hello cannot be converted to float64 type
	assert.False(t, ok)
}
