package advanced_types

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArraysLiterals(t *testing.T) {
	t.Run("declaration and initialization", func(t *testing.T) {
		var arr1 = [5]int{18, 20, 15, 22, 16}
		var arr2 [5]int
		arr2 = [5]int{18, 20, 15, 22, 16}
		assert.Equal(t, arr1, arr2)
	})

	t.Run("[...] and []", func(t *testing.T) {
		// var emptyArr [...]int --> use of [...] array outside of array literal
		var arr = [...]int{5, 6, 7, 8, 22} // array
		assert.Equal(t, "[5]int", fmt.Sprint(reflect.TypeOf(arr)))
		var sli = []int{5, 6, 7, 8, 22} // slice
		assert.Equal(t, "[]int", fmt.Sprint(reflect.TypeOf(sli)))
		assert.Equal(t, cap(sli), 5)
	})

	t.Run("skip indexes", func(t *testing.T) {
		var arr = [5]string{3: "Chris", 4: "Ron"} // array
		var sli = []string{3: "Chris", 4: "Ron"}  // slice

		assert.Equal(t, 5, cap(arr))
		assert.Equal(t, 5, len(arr))
		assert.Equal(t, 5, cap(sli))
		assert.Equal(t, 5, len(sli))
	})

	// summary: [count] and [...] means array while [] means slice

	t.Run("empty interface", func(t *testing.T) {
		var anyType = [...]interface{}{"1", 2, true}
		assert.Equal(t, 3, cap(anyType))
	})

	t.Run("new keyword", func(t *testing.T) {
		var arr1 = new([5]int)

		for _, v := range arr1 {
			assert.Equal(t, 0, v)
		}

		for _, v := range *arr1 {
			assert.Equal(t, 0, v)
		}

		arr2 := *arr1 // copied value
		arr2[2] = 5
		assert.Equal(t, 0, arr1[2]) // not modified

		var arr3 = arr1
		arr3[2] = 5
		assert.Equal(t, 5, arr1[2])
	})
}

func TestArraysIteration(t *testing.T) {
	var intArr [5]int
	for i := 0; i < len(intArr); i++ {
		assert.Equal(t, 0, intArr[i])
		intArr[i] = i * 2
	}

	for i, v := range intArr {
		assert.Equal(t, v, intArr[i])
	}

	t.Run("nested array", func(t *testing.T) {
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
	})
}

func TestArraysCapMethod(t *testing.T) {
	var arr1 = [...]int{5, 6, 7, 8, 22} // array
	var arr2 = [5]int{5, 6, 7, 8, 22}   // array
	var sli = []int{5, 6, 7, 8, 22}     // slice
	assert.Equal(t, 5, cap(arr1))
	assert.Equal(t, 5, cap(arr2))
	assert.Equal(t, 5, cap(sli))
}

func TestArrayAsParameters(t *testing.T) {
	t.Run("array vs pointer to array", func(t *testing.T) {
		// accept [3]int
		funcAcceptArray := func(arr [3]int) string {
			arr[0] = 1
			return fmt.Sprint(arr)
		}

		// accept pointer to [3]int
		funcAcceptArrayPointer := func(arr *[3]int) string {
			arr[0] = 2
			return fmt.Sprint(arr)
		}

		var arr [3]int

		assert.Equal(t, "[1 0 0]", funcAcceptArray(arr))
		assert.Equal(t, [3]int{0, 0, 0}, arr)

		assert.Equal(t, "&[2 0 0]", funcAcceptArrayPointer(&arr))
		assert.Equal(t, [3]int{2, 0, 0}, arr)
	})
}

func TestArraysPointerOrSlicesAsArguments(t *testing.T) {
	// pass pointers to array
	Sum := func(numbers *[3]float64) (sum float64) {
		for _, v := range numbers {
			sum += v
		}

		return
	}

	// pass slices
	SliceSum := func(numbers []float64) (sum float64) {
		for _, v := range numbers {
			sum += v
		}

		return
	}
	array := [3]float64{1.1, 2.2, 3.3}
	assert.Equal(t, 6.6, Sum(&array))
	slice := array[:]
	assert.Equal(t, 6.6, SliceSum(slice))
}
