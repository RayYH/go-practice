package variables

import (
	"fmt"
	"reflect"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGlobalVariablesInitializationAndScope(t *testing.T) {
	t.Run("global variables are visible to other files", func(t *testing.T) {
		assert.Equal(t, "This is a string.", globalString)
		assert.Equal(t, "This is also a string.", GlobalString)
	})

	t.Run("global variables can be modified", func(t *testing.T) {
		assert.Equal(t, 24, myAge)
		assert.Equal(t, "Ray", myName)
		myAge = 25
		myName = "Ray Hong"
		assert.Equal(t, 25, myAge)
		assert.Equal(t, "Ray Hong", myName)
	})

	t.Run("home and user are not empty", func(t *testing.T) {
		assert.NotNil(t, HOME)
		assert.NotNil(t, USER)
	})
}

func TestContentOfVariablesCanBeModifiedThroughPointers(t *testing.T) {
	var os = runtime.GOOS
	assert.NotNil(t, os)

	// q == p --> os (p and q are both pointers)
	// pointers holds the references to the variables
	var p = &os
	var q = p
	assert.Equal(t, p, q)
	assert.Equal(t, "*string", fmt.Sprint(reflect.TypeOf(p)))
	assert.Equal(t, *p, *q)
	assert.Equal(t, "string", fmt.Sprint(reflect.TypeOf(*q)))

	*p = "new string"
	assert.Equal(t, os, "new string")
}

func TestVariablesDeclarationAndInitialization(t *testing.T) {
	t.Run("declare first, then initialize", func(t *testing.T) {
		// var declares 1 or more variables
		var a, b int
		var c string
		a, b, c = 5, 7, "abc"
		assert.Equal(t, 5, a)
		assert.Equal(t, 7, b)
		assert.Equal(t, "abc", c)
	})

	t.Run("declaration and initialization at the same time", func(t *testing.T) {
		var v = 10
		assert.Equal(t, v, 10)
	})

	t.Run(":= short assignment statement", func(t *testing.T) {
		d, e, f := 5, 7, "abc"
		assert.Equal(t, 5, d)
		assert.Equal(t, 7, e)
		assert.Equal(t, "abc", f)

		// though g has been declared, h and i were not declared, so we can use := syntax
		var g int
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

	t.Run("Variables declared without a corresponding initialization are zero-valued.", func(t *testing.T) {
		var v1 int
		var v2 float64
		var v3 string
		var v4 []int
		assert.Equal(t, v1, 0)
		assert.Equal(t, v2, 0.0)
		assert.Equal(t, v3, "")
		assert.Nil(t, v4)
	})

	t.Run("Advanced variables declaration", func(t *testing.T) {
		var v5 struct {
			f int
		} // struct
		var v6 struct{}        // empty struct
		var v7 *int            // pointer to int type
		var v8 map[string]int  // map
		var v9 func(a int) int // func
		assert.NotNil(t, v5)
		assert.NotNil(t, v6)
		assert.Zero(t, v5.f)
		assert.Nil(t, v7)
		assert.Nil(t, v8)
		assert.Nil(t, v9)
	})
}

func TestVariableInitializedInsideInitFunc(t *testing.T) {
	assert.Equal(t, 0.7853981633974483, declaredVariable)
}

// anonymous variables
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
