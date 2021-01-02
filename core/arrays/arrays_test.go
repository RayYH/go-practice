package arrays

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArraysIteration(t *testing.T) {
	// declare an array using var keyword
	var intArr [5]int

	// use [] to access array elements
	for i := 0; i < 5; i++ {
		assert.Equal(t, intArr[i], 0)
	}

	// use [] to modify array elements
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

func TestArraysLiterals(t *testing.T) {
	// declare an array of the specified length and initialize it
	var arr1 = [5]int{18, 20, 15, 22, 16}

	// declare first, then initialize
	var arr2 [5]int
	arr2 = [5]int{18, 20, 15, 22, 16}

	assert.Equal(t, arr1, arr2)

	// arr3 is array, arr4 is slices
	var arr3 = [...]int{5, 6, 7, 8, 22}
	var arr4 = []int{5, 6, 7, 8, 22}

	assert.Equal(t, 5, cap(arr3))
	assert.Equal(t, 5, len(arr3))
	assert.Equal(t, 5, cap(arr4))
	assert.Equal(t, 5, len(arr4))
	assert.NotEqual(t, arr3, arr4)

	// only indexes 3 and 4 are assigned values
	// other elements are set to zero value (empty strings), and the array length is 5
	var arr5 = [5]string{3: "Chris", 4: "Ron"} // array
	var arr6 = []string{3: "Chris", 4: "Ron"}  // slice

	assert.Equal(t, 5, cap(arr5))
	assert.Equal(t, 5, len(arr5))
	assert.Equal(t, 5, cap(arr6))
	assert.Equal(t, 5, len(arr6))
	assert.NotEqual(t, arr5, arr6)
	assert.Equal(t, arr5[0], "")
}

func TestArraysOrSlicesAsArguments(t *testing.T) {
	Sum := func(numbers *[3]float64) (sum float64) {
		for _, v := range numbers {
			sum += v
		}

		return
	}

	SliceSum := func(numbers []float64) (sum float64) {
		for _, v := range numbers {
			sum += v
		}

		return
	}
	array := [3]float64{1.1, 2.2, 3.3}
	assert.Equal(t, 6.6, Sum(&array))
	slice := []float64{1.1, 2.2, 3.3}
	assert.Equal(t, 6.6, SliceSum(slice))
}

func TestArraysFormattedAsString(t *testing.T) {
	// The parameter is an array of [3]int type
	funcAcceptArray := func(arr [3]int) string {
		arr[0] = 1
		return fmt.Sprint(arr)
	}

	// The parameter is a pointer to an array of [3]int type
	funcAcceptArrayPointer := func(arr *[3]int) string {
		arr[0] = 2
		return fmt.Sprint(arr)
	}

	// declare an array
	var arr [3]int
	// passes a copy of arr
	assert.Equal(t, "[1 0 0]", funcAcceptArray(arr))
	// arr was not modified
	assert.Equal(t, [3]int{0, 0, 0}, arr)
	// passes a pointer to arr
	assert.Equal(t, "&[2 0 0]", funcAcceptArrayPointer(&arr))
	// arr was modified
	assert.Equal(t, [3]int{2, 0, 0}, arr)
}

func TestArraysCanBeNested(t *testing.T) {
	type pixel int

	const (
		WIDTH  = 192
		HEIGHT = 108
	)

	var screen [WIDTH][HEIGHT]pixel

	for y := 0; y < HEIGHT; y++ {
		for x := 0; x < WIDTH; x++ {
			// nested loop to access elements of two-dimensional array
			screen[x][y] = 0
		}
	}

	// use for-range syntax to access elements
	for row := range screen {
		for column := range screen[row] {
			assert.Equal(t, pixel(0), screen[row][column], "the pixel value should be the zero value 0")
		}
	}
}
