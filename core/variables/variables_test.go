package variables

import (
	"fmt"
	"reflect"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeclarationAndInitialization(t *testing.T) {
	t.Run("declare first, then initialize", func(t *testing.T) {
		var a, b int
		var c string
		a, b, c = 5, 7, "abc"
		assert.Equal(t, 5, a)
		assert.Equal(t, 7, b)
		assert.Equal(t, "abc", c)
	})

	t.Run("declaration and initialization at the same time", func(t *testing.T) {
		var v = 10
		assert.Equal(t, 10, v)
	})

	t.Run(":= short assignment", func(t *testing.T) {
		d, e, f := 5, 7, "abc"
		assert.Equal(t, 5, d)
		assert.Equal(t, 7, e)
		assert.Equal(t, "abc", f)

		var g, h int
		// i is not declared
		g, h, i := 1, 2, 3
		assert.Equal(t, 1, g)
		assert.Equal(t, 2, h)
		assert.Equal(t, 3, i)
	})

	t.Run("swap two variables' value", func(t *testing.T) {
		a := 5
		b := 7
		a, b = b, a
		assert.Equal(t, 7, a)
		assert.Equal(t, 5, b)
	})

	t.Run("variables declared without a corresponding initialization are zero-valued", func(t *testing.T) {
		var v1 int
		var v2 float64
		var v3 string
		assert.Equal(t, 0, v1)
		assert.Equal(t, 0.0, v2)
		assert.Equal(t, "", v3)
	})

	t.Run("declaration of variables holding advance types", func(t *testing.T) {
		var v4 []int
		var v5 struct { // struct
			f int
		}
		var v6 struct{}        // empty struct
		var v7 *int            // pointer to int type
		var v8 map[string]int  // map
		var v9 func(a int) int // func
		assert.Nil(t, v4)
		assert.NotNil(t, v5)
		assert.NotNil(t, v6)
		assert.Zero(t, v5.f)
		assert.Nil(t, v7)
		assert.Nil(t, v8)
		assert.Nil(t, v9)
	})
}

func TestGlobalVariablesInitializationAndScope(t *testing.T) {
	t.Run("global variables are visible to other files", func(t *testing.T) {
		assert.Equal(t, "This is a string.", globalString)
		assert.Equal(t, "This is also a string.", GlobalString)
	})

	t.Run("global variables declared but no initialization have zero values", func(t *testing.T) {
		assert.Equal(t, "", emptyGlobalVar)
	})

	t.Run("global variables can be modified", func(t *testing.T) {
		assert.Equal(t, 24, myAge)
		assert.Equal(t, "Ray", myName)
		myAge, myName = 25, "Ray Hong"
		assert.Equal(t, 25, myAge)
		assert.Equal(t, "Ray Hong", myName)
	})

	t.Run("global variables can be initialized inside init func", func(t *testing.T) {
		assert.Equal(t, 0.7853981633974483, declaredVariable)
	})
}

func TestUsingBlankIdentifierToDiscardValues(t *testing.T) {
	var _, age = "Ray", 24
	assert.Equal(t, 24, age)

	getName := func() (firstName, middleName, lastName string) {
		firstName, middleName, lastName = "Ray", "Young", "Hong"
		return
	}

	_, middle, _ := getName()
	assert.Equal(t, "Young", middle)
}

func TestContentOfVariablesCanBeModifiedThroughPointers(t *testing.T) {
	var os = runtime.GOOS
	assert.NotNil(t, os)

	var p = &os
	// only reference (location) copied when assigned
	var q = p

	// both p and q are *string type (a pointer to a string)
	assert.Equal(t, p, q)
	assert.Equal(t, "*string", fmt.Sprint(reflect.TypeOf(p)))
	// both *p and *q are string type
	assert.Equal(t, *p, *q)
	assert.Equal(t, "string", fmt.Sprint(reflect.TypeOf(*q)))

	// if variable's value has been changed, then all references to this variable will refer to the new value
	*p = "new string"
	assert.Equal(t, os, "new string")
	assert.Equal(t, p, q)
	assert.Equal(t, *p, *q)
}
