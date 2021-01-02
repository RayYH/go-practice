package variables

import (
	"github.com/stretchr/testify/assert"
	"runtime"
	"testing"
)

func TestGlobalString(t *testing.T) {
	assert.Equal(t, "This is a string.", globalString)
	assert.Equal(t, 24, myAge)
	assert.Equal(t, "Ray", myName)
	myAge = 25
	assert.Equal(t, 25, myAge)
}

func TestHomeAndUserNotEmpty(t *testing.T) {
	assert.NotNil(t, HOME)
	assert.NotNil(t, USER)
}

func TestModifyingValueViaPointers(t *testing.T) {
	var os = runtime.GOOS
	assert.NotNil(t, os)

	var p = &os
	var q = p
	assert.Equal(t, p, q)
	assert.Equal(t, *p, *q)

	*p = "new string"
	assert.Equal(t, os, *p)
}

func TestLocalVariablesOne(t *testing.T) {
	// declare first, then initialize
	var a, b int
	var c string
	a, b, c = 5, 7, "abc"
	assert.Equal(t, 5, a)
	assert.Equal(t, 7, b)
	assert.Equal(t, "abc", c)

	// swap a and b
	a, b = b, a
	assert.Equal(t, 7, a)
	assert.Equal(t, 5, b)

	// use := syntax, d, e, f has not been declared yet
	d, e, f := 5, 7, "abc"
	assert.Equal(t, 5, d)
	assert.Equal(t, 7, e)
	assert.Equal(t, "abc", f)

	// though g has been declared, but h and i were not declared
	// so we can use := syntax
	var g int
	g, h, i := 1, 2, 3
	assert.Equal(t, 1, g)
	assert.Equal(t, 2, h)
	assert.Equal(t, 3, i)
}

func TestVariableInitializedInsideInitFunc(t *testing.T) {
	assert.Equal(t, 0.7853981633974483, declaredVariable)
}

func TestUsingBlankIdentifierToDiscardValues(t *testing.T) {
	var _, age = "Ray", 24
	assert.Equal(t, 24, age)
	getName := func() (firstName, middleName, lastName string) {
		firstName = "Ray"
		middleName = "Young"
		lastName = "Hong"

		return
	}

	_, middle, _ := getName()
	assert.Equal(t, "Young", middle)
}
