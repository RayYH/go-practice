package functions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFunctionDeclaration(t *testing.T) {
	multiply3Nums := func(a, b, c int) int {
		return a * b * c
	}

	assert.Equal(t, 60, multiply3Nums(2, 5, 6))
}

func TestFunctionCanReturnMultiValues(t *testing.T) {
	getX2AndX3 := func(n int) (int, int) {
		return n * 2, n * 3
	}

	x2, x3 := getX2AndX3(2)
	assert.Equal(t, 4, x2)
	assert.Equal(t, 6, x3)
}

func TestReturnValuesCanBeNamed(t *testing.T) {
	// Go's return values may be named.
	// If so, they are treated as variables defined at the top of the function.
	getNamedX2AndX3 := func(n int) (x2, x3 int) {
		x2 = n * 2
		x3 = n * 3
		// A return statement without arguments returns the named return values
		// Naked return statements should be used only in short functions
		return
	}

	x2, x3 := getNamedX2AndX3(2)
	assert.Equal(t, 4, x2)
	assert.Equal(t, 6, x3)
}

func TestParamsCanBeModifiedInsideFuncByPassingRefs(t *testing.T) {
	// reply can be modified inside this func
	multiply := func(a, b int, reply *int) {
		*reply = a * b
	}

	n := 0
	reply := &n
	multiply(3, 4, reply)
	assert.Equal(t, 12, *reply)
}

func TestRestParameters(t *testing.T) {
	// rest parameters
	min := func(s ...int) int {
		if len(s) == 0 {
			return 0
		}

		minValue := s[0]

		for _, v := range s {
			if v < minValue {
				minValue = v
			}
		}

		return minValue
	}

	assert.Equal(t, 1, min(1, 2, 3, 4, 5))
	assert.Equal(t, 1, min([]int{1, 2, 3, 4, 5}...))
	assert.Equal(t, 2, min([]int{1, 2, 3, 4, 5}[1:]...))
	assert.Equal(t, 0, min())
	assert.Equal(t, -5, min(1, 2, -5, 3, 111))
}

func TestRecursion(t *testing.T) {
	result := 0
	nums := []int{1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89}
	for i := 0; i <= 10; i++ {
		result = fibonacci(i)
		assert.Equal(t, nums[i], result)
	}
}

func TestEmptyInterfaceCanAcceptParamsOfAnyTypes(t *testing.T) {
	getType := func(args ...interface{}) string {
		for _, arg := range args {
			switch arg.(type) {
			case int:
				return "int"
			case string:
				return "string"
			case int64:
				return "int64"
			default:
				return "unknown"
			}
		}
		return ""
	}

	var v1 = 1 // inferred int
	var v2 int64 = 234
	var v3 = "Hello" // inferred string
	var v4 float32 = 1.234
	assert.Equal(t, getType(v1), "int")
	assert.Equal(t, getType(v2), "int64")
	assert.Equal(t, getType(v3), "string")
	assert.Equal(t, getType(v4), "unknown")
}
