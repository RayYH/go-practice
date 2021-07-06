package basic_types

import (
	"fmt"
	"math"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTypeConversions(t *testing.T) {
	var n int16 = 34
	var m int32
	m = int32(n)
	assert.NotEqual(t, reflect.TypeOf(n), reflect.TypeOf(m))
	assert.Equal(t, fmt.Sprint(m), fmt.Sprint(n))

	var x, y = 3, 4
	var f = math.Sqrt(float64(x*x + y*y))
	var z = uint(f)
	var r uint = 5
	assert.Equal(t, r, z)

	t.Run("lost precision", func(t *testing.T) {
		f64 := 3.1415926535897932384626433
		f32 := float32(f64)
		assert.Equal(t, fmt.Sprintf("%v %v", f32, f64), "3.1415927 3.141592653589793")
	})
}

func TestTypeAliases(t *testing.T) {
	// = means alias
	type (
		A1 = string
		A2 = A1
	)

	// not using = mean type definition
	type (
		B1 string
		B2 B1
		B3 []B1
		B4 B3
	)

	var a1 A1
	var a2 A2
	var b1 B1
	var b2 B2
	var b3 B3
	var b4 B4

	assert.Equal(t, "", a1)
	assert.Equal(t, "", a2)
	assert.Equal(t, B1(""), b1)
	assert.Equal(t, B2(""), b2)
	assert.Nil(t, b3)
	assert.Nil(t, b4)
}
