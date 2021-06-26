package advanced_types

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArraysIteration(t *testing.T) {
	// declare an array using var keyword
	// each element in intArr has the zero value (0) of int type
	var intArr [5]int

	// we can use C-style for loop to iterate an array
	// len(intArr) will return the length of given array
	for i := 0; i < len(intArr); i++ {
		// we can access and modify elements via [] syntax
		assert.Equal(t, intArr[i], 0)
		intArr[i] = i * 2
	}

	// use for range to access array elements (key-value pairs)
	for i, v := range intArr {
		assert.Equal(t, v, intArr[i])
	}
}

func TestArraysLiterals(t *testing.T) {
	// declare an array of the specified length and initialize it
	var arr1 = [5]int{18, 20, 15, 22, 16}

	// separate declaration and initialization
	var arr2 [5]int
	arr2 = [5]int{18, 20, 15, 22, 16}

	assert.Equal(t, arr1, arr2)

	// arr3 is an array, sli3 is a slice
	// when use ..., the compiler will count the array elements for you
	var arr3 = [...]int{5, 6, 7, 8, 22} // array
	var sli3 = []int{5, 6, 7, 8, 22}    // slice
	assert.NotEqual(t, arr3, sli3)

	// only indexes 3 and 4 are assigned values
	// other elements are set to zero value "" (empty strings)
	// the array length is 5
	var arr4 = [5]string{3: "Chris", 4: "Ron"} // array
	var sli4 = []string{3: "Chris", 4: "Ron"}  // slice

	assert.Equal(t, 5, cap(arr4))
	assert.Equal(t, 5, len(arr4))
	assert.Equal(t, 5, cap(sli4))
	assert.Equal(t, 5, len(sli4))
	assert.NotEqual(t, arr4, sli4)
	// summary, [num] or [...] means arrays while [] means slices
}

func TestArraysCapMethod(t *testing.T) {
	// cap tells you the capacity of the underlying array
	// so both slices and arrays can call cap method
	var arr1 = [...]int{5, 6, 7, 8, 22} // array
	var arr2 = [5]int{5, 6, 7, 8, 22}   // array
	var sli = []int{5, 6, 7, 8, 22}     // slice
	assert.Equal(t, 5, cap(arr1))
	assert.Equal(t, 5, cap(arr2))
	assert.Equal(t, 5, cap(sli))
}

func TestArraysPointerOrSlicesAsArguments(t *testing.T) {
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

	var arr [3]int

	// passes a copy of arr, arr has not been modified
	assert.Equal(t, "[1 0 0]", funcAcceptArray(arr))
	assert.Equal(t, [3]int{0, 0, 0}, arr)

	// passes a pointer to arr, arr has been modified
	assert.Equal(t, "&[2 0 0]", funcAcceptArrayPointer(&arr))
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
			screen[x][y] = 1
		}
	}

	// use for-range syntax to access elements
	for row := range screen {
		for column := range screen[row] {
			assert.Equal(t, pixel(1), screen[row][column])
		}
	}
}

func ExampleDisplayArrays() {
	DisplayArrays()
	// Output: '\x00'
	// '\x00'
	// 0 0
	// 0 0
	// 0 0
	// 0 0
	// 0.00
	// 1.11
	// 2.22
	// 1234
}
