package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIterateAnArray(t *testing.T) {
	var intArr [5]int
	for i := 0; i < 5; i++ {
		assert.Equal(t, intArr[i], 0, "default value of int type is 0")
	}
	for i := 0; i < 5; i++ {
		intArr[i] = i * 2
	}
	// use i < len(intArr) as condition
	for i := 0; i < len(intArr); i++ {
		assert.Equal(t, intArr[i], i*2)
	}
	// use for range
	for i, v := range intArr {
		assert.Equal(t, v, intArr[i])
	}
}

func ExampleArrayLiterals() {
	var arr1 = [5]int{18, 20, 15, 22, 16}
	var arr2 = [...]int{5, 6, 7, 8, 22}
	var arr3 = []int{5, 6, 7, 8, 22}
	var arr4 = [5]string{3: "Chris", 4: "Ron"}
	var arr5 = []string{3: "Chris", 4: "Ron"}
	fmt.Println(arr1, arr2, arr3, arr4, arr5)
	// Output:
	// [18 20 15 22 16] [5 6 7 8 22] [5 6 7 8 22] [   Chris Ron] [   Chris Ron]
}

func TestSum(t *testing.T) {
	// here must use [3]float64, cannot be []float64
	array := [3]float64{1.1, 2.2, 3.3}
	assert.Equal(t, Sum(&array), 6.6)
}
