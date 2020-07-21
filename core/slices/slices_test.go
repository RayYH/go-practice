package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func ExampleInitializeSlice() {
	names := []string{"leo", "jessica", "paul"}
	checks := make([]bool, 10)
	scores := make([]int, 2, 20)
	numbers := new([20]int)[0:2]
	fmt.Println(names)
	fmt.Println(checks)
	fmt.Println(scores)
	fmt.Println(numbers)
	// Output:
	// [leo jessica paul]
	// [false false false false false false false false false false]
	// [0 0]
	// [0 0]
}

func TestSliceUsage(t *testing.T) {
	var arr [6]int
	var slice = arr[2:5]

	for i := 0; i < len(arr); i++ {
		arr[i] = i * 2
	}

	for i := 0; i < len(slice); i++ {
		assert.Equal(t, slice[i], (i+2)*2)
	}

	// 切片的长度就是它所包含的元素个数。
	// 切片的容量是从它的第一个元素开始数，到其底层数组元素末尾的个数
	assert.Equal(t, len(arr), 6)
	assert.Equal(t, len(slice), 3)
	assert.Equal(t, cap(slice), 4)
}
