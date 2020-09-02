package arrays

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIterateAnArray(t *testing.T) {
	var intArr [5]int

	// use [] syntax to access array elements
	for i := 0; i < 5; i++ {
		assert.Equal(t, intArr[i], 0, "default value of int type is 0")
	}
	// use [] syntax to modify array elements
	for i := 0; i < 5; i++ {
		intArr[i] = i * 2
	}
	// len(intArr) will return the length of given array
	for i := 0; i < len(intArr); i++ {
		assert.Equal(t, intArr[i], i*2)
	}
	// use for range to access array elements
	for i, v := range intArr {
		assert.Equal(t, v, intArr[i])
	}
}

func TestArrayLiterals(t *testing.T) {
	// Declare an array of the specified length
	// declaration + initialization
	var arr1 = [5]int{18, 20, 15, 22, 16}
	// declare first, then initialize
	var arr2 [5]int
	arr2 = [5]int{18, 20, 15, 22, 16}

	// Do not declare the length, the ... can be ignored
	// they actually become slices
	var arr3 = [...]int{5, 6, 7, 8, 22}
	var arr4 = []int{5, 6, 7, 8, 22}
	assert.Equal(t, 5, cap(arr3))
	assert.Equal(t, 5, len(arr3))
	assert.Equal(t, 5, cap(arr4))
	assert.Equal(t, 5, len(arr4))

	// Only indexes 3 and 4 are assigned values, other elements are set to empty strings, and the array length is 5
	var arr5 = [5]string{3: "Chris", 4: "Ron"}
	var arr6 = []string{3: "Chris", 4: "Ron"}
	assert.Equal(t,
		"[18 20 15 22 16] [18 20 15 22 16] [5 6 7 8 22] [5 6 7 8 22] [   Chris Ron] [   Chris Ron]",
		fmt.Sprint(arr1, arr2, arr3, arr4, arr5, arr6))
}

func TestSum(t *testing.T) {
	// here must be [3]float64, cannot be []float64 (of slice type)
	array := [3]float64{1.1, 2.2, 3.3}
	assert.Equal(t, 6.6, Sum(&array), "Sum of [1.1, 2.2, 3.3] should be 6.6")
	slice := []float64{1.1, 2.2, 3.3}
	assert.Equal(t, 6.6, SliceSum(slice), "Sum of [1.1, 2.2, 3.3] should be 6.6")
}

// The parameter is an array [3]int
func f(a [3]int) string {
	a[0] = 1
	return fmt.Sprint(a)
}

// The parameter is a pointer to an array [3]int
func fp(a *[3]int) string {
	a[0] = 2
	return fmt.Sprint(a)
}

func TestOutputArrayType(t *testing.T) {
	var ar [3]int
	assert.Equal(t, "[1 0 0]", f(ar))    // passes a copy of ar
	assert.Equal(t, [3]int{0, 0, 0}, ar) // ar was not modified
	assert.Equal(t, "&[2 0 0]", fp(&ar)) // passes a pointer to ar
	assert.Equal(t, [3]int{2, 0, 0}, ar) // ar was modified
}

// 10-times small for testing consideration
const (
	WIDTH  = 192
	HEIGHT = 108
)

type pixel int

var screen [WIDTH][HEIGHT]pixel

func TestMultipleArray(t *testing.T) {
	for y := 0; y < HEIGHT; y++ {
		for x := 0; x < WIDTH; x++ {
			// nested loop to access elements of two-dimensional array
			screen[x][y] = 0
		}
	}

	// use for-range syntax to access elements
	for row := range screen {
		for column := range screen[row] {
			assert.Equal(t, pixel(0), screen[row][column], "the pixel value should be 0")
		}
	}
}
