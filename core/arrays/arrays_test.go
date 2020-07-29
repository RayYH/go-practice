package main

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
	// use i < len(intArr) as condition
	for i := 0; i < len(intArr); i++ {
		assert.Equal(t, intArr[i], i*2)
	}
	// use for range to access array elements
	for i, v := range intArr {
		assert.Equal(t, v, intArr[i])
	}
}

func ExampleArrayLiterals() {
	// 声明指定长度的数组
	var arr1 = [5]int{18, 20, 15, 22, 16}
	// 不声明长度，下面的 ... 可以忽略, 从技术上说它们其实变成了切片
	var arr2 = [...]int{5, 6, 7, 8, 22}
	var arr3 = []int{5, 6, 7, 8, 22}
	// 只有索引 3 和 4 被赋予实际的值，其他元素都被设置为空的字符串，数组长度为 5
	var arr4 = [5]string{3: "Chris", 4: "Ron"}
	var arr5 = []string{3: "Chris", 4: "Ron"}
	fmt.Println(arr1, arr2, arr3, arr4, arr5)
	// Output:
	// [18 20 15 22 16] [5 6 7 8 22] [5 6 7 8 22] [   Chris Ron] [   Chris Ron]
}

func TestSum(t *testing.T) {
	// here must be [3]float64, cannot be []float64 (of slice type)
	array := [3]float64{1.1, 2.2, 3.3}
	assert.Equal(t, 6.6, Sum(&array), "Sum of [1.1, 2.2, 3.3] should be 6.6")
	slice := []float64{1.1, 2.2, 3.3}
	assert.Equal(t, 6.6, SliceSum(slice), "Sum of [1.1, 2.2, 3.3] should be 6.6")
}

// 参数为一个数组
func f(a [3]int) { fmt.Println(a) }

// 参数为一个数组指针
func fp(a *[3]int) { fmt.Println(a) }

func ExampleOutputArrayType() {
	var ar [3]int
	f(ar)   // passes a copy of ar
	fp(&ar) // passes a pointer to ar
	// Output:
	// [0 0 0]
	//&[0 0 0]
}

// 10-times small for testing
const (
	WIDTH  = 192
	HEIGHT = 108
)

type pixel int

var screen [WIDTH][HEIGHT]pixel

func TestMultipleArray(t *testing.T) {
	for y := 0; y < HEIGHT; y++ {
		for x := 0; x < WIDTH; x++ {
			screen[x][y] = 0
		}
	}

	// or use for-range syntax to access elements
	for row := range screen {
		for column := range screen[row] {
			assert.Equal(t, pixel(0), screen[row][column], "the pixel value should be 0")
		}
	}
}
