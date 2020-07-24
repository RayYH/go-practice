package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func ExampleBasicSliceOutput() {
	b := []byte{'g', 'o', 'l', 'a', 'n', 'g'}
	fmt.Println(b[1:4])
	fmt.Println(b[:2])
	fmt.Println(b[2:])
	fmt.Println(b[:])
	// Output:
	// [111 108 97]
	// [103 111]
	// [108 97 110 103]
	// [103 111 108 97 110 103]
}

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

// new() 和 make() 的区别 - new 函数分配内存，make 函数初始化
// new(T) 为每个新的类型 T 分配一片内存，初始化为 0 并且返回类型为 *T 的内存地址，
// 这种方法返回一个指向类型为 T，值为 0 的地址的指针，它适用于值类型如数组和结构体，相当于 &T{}
// make(T) 返回一个类型为 T 的初始值，它只适用于 3 种内建的引用类型: slice, map, channel.

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

func TestSlicesShareCommonData(t *testing.T) {
	// 多个切片如果表示同一个数组的片段，它们可以共享数据
	// 因此一个切片和相关数组的其他切片是共享存储的
	// 不同的数组总是代表不同的存储
	// 数组实际上是切片的构建块
	values := []int{1, 2, 3, 4, 5, 6}
	slice1 := values[1:]
	slice2 := values[1:]
	assert.Equal(t, slice1[0], 2)
	assert.Equal(t, slice2[0], 2)
	values[1] = 9
	assert.Equal(t, values[1], 9)
	assert.Equal(t, slice1[0], 9)
	assert.Equal(t, slice2[0], 9)
}

func TestSliceMovement(t *testing.T) {
	array := []int{1, 2, 3, 4, 5, 6}
	slice := array[:]
	slice = slice[1:]
	assert.Equal(t, len(slice), 5)
}

func ExampleSliceForRange() {
	var slice = make([]int, 4)
	slice[0] = 1
	slice[1] = 2
	slice[2] = 3
	slice[3] = 4
	for ix, value := range slice {
		fmt.Printf("Slice at %d is: %d\n", ix, value)
	}
	// Output:
	// Slice at 0 is: 1
	// Slice at 1 is: 2
	// Slice at 2 is: 3
	// Slice at 3 is: 4
}

func ExampleSliceForRange2() {
	seasons := []string{"Spring", "Summer", "Autumn", "Winter"}

	// element
	for ix, season := range seasons {
		fmt.Printf("Season %d is %s\n", ix, season)
	}

	// only value
	var season string
	for _, season = range seasons {
		fmt.Printf("%s\n", season)
	}

	// only index
	for ix := range seasons {
		fmt.Printf("%d\n", ix)
	}

	// Output:
	// Season 0 is Spring
	// Season 1 is Summer
	// Season 2 is Autumn
	// Season 3 is Winter
	// Spring
	// Summer
	// Autumn
	// Winter
	// 0
	// 1
	// 2
	// 3
}

func ExampleReSlice() {
	capNum := 10
	slice := make([]int, 0, capNum)

	for i := 0; i < cap(slice); i++ {
		slice = slice[0 : i+1]
		slice[i] = i
		fmt.Printf("The length of slice is %d\n", len(slice))
	}

	for i := 0; i < len(slice); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice[i])
	}
	// Output:
	// The length of slice is 1
	// The length of slice is 2
	// The length of slice is 3
	// The length of slice is 4
	// The length of slice is 5
	// The length of slice is 6
	// The length of slice is 7
	// The length of slice is 8
	// The length of slice is 9
	// The length of slice is 10
	// Slice at 0 is 0
	// Slice at 1 is 1
	// Slice at 2 is 2
	// Slice at 3 is 3
	// Slice at 4 is 4
	// Slice at 5 is 5
	// Slice at 6 is 6
	// Slice at 7 is 7
	// Slice at 8 is 8
	// Slice at 9 is 9
}

func TestReSlice(t *testing.T) {
	var ar = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	var a = ar[5:7] // reference to subarray {5,6} - len(a) is 2 and cap(a) is 5
	assert.Equal(t, a, []int{5, 6})
	assert.Equal(t, 2, len(a))
	assert.Equal(t, 5, cap(a))
	a = a[0:4] // ref of subarray {5,6,7,8} - len(a) is now 4 but cap(a) is still 5
	assert.Equal(t, a, []int{5, 6, 7, 8})
	assert.Equal(t, 4, len(a))
	assert.Equal(t, 5, cap(a))

	assert.Equal(t, 0, len(a[1:1]))
	assert.Equal(t, 1, len(a[1:2]))
}

func TestDuplication(t *testing.T) {
	slFrom := []int{1, 2, 3}
	slTo := make([]int, 10)
	n := copy(slTo, slFrom) // n = copied elements
	assert.Equal(t, 3, n)
	assert.Equal(t, slTo, []int{1, 2, 3, 0, 0, 0, 0, 0, 0, 0})
	slice := []int{1, 2, 3}
	slice = append(slice, 4, 5, 6)
	assert.Equal(t, slice, []int{1, 2, 3, 4, 5, 6})
}

func ExampleGenerateSliceFromString() {
	s := "\u00ff\u754c"
	for i, c := range s {
		fmt.Printf("%d:%c ", i, c)
	}
	// Output:
	// 0:ÿ 2:界
}

func TestAppendStringToCharArray(t *testing.T) {
	var b []byte
	b = []byte{'H', 'e', 'l', 'l', 'o'}
	var s string
	s = " World"
	b = append(b, s...)
	assert.Equal(t, string(b), "Hello World")
	assert.Equal(t, b, []byte{'H', 'e', 'l', 'l', 'o', ' ', 'W', 'o', 'r', 'l', 'd'})
}

// 在内存中，一个字符串实际上是一个双字结构，即一个指向实际数据的指针和记录字符串长度的整数
// 因为指针对用户来说是完全不可见，因此我们可以依旧把字符串看做是一个值类型，也就是一个字符数组
// Go 语言中的字符串是不可变的
func TestStringIsImmutable(t *testing.T) {
	s := "Hello"
	c := []byte(s)
	c[0] = 'c'
	assert.Equal(t, string(c), "cello")
}

func TestCompare(t *testing.T) {
	assert.Equal(t, Compare([]byte{1, 2, 3}, []byte{1, 2, 3}), 0)
	assert.Equal(t, Compare([]byte{1, 2, 4}, []byte{1, 2, 3}), 1)
	assert.Equal(t, Compare([]byte{1, 2, 3}, []byte{1, 2, 4}), -1)
	assert.Equal(t, Compare([]byte{1, 2, 3, 4}, []byte{1, 2, 3}), 1)
	assert.Equal(t, Compare([]byte{1, 2, 3}, []byte{1, 2, 3, 4}), -1)
}

func TestAppendMethod(t *testing.T) {
	var i, j int
	var x int
	a := []int{1, 2, 3}
	b := []int{4, 5, 6}
	// 将切片 b 的元素追加到切片 a
	a = append(a, b...)
	assert.Equal(t, a, []int{1, 2, 3, 4, 5, 6})
	// 复制切片 a 的元素到新的切片 b
	b = make([]int, len(a))
	copy(b, a)
	assert.Equal(t, b, []int{1, 2, 3, 4, 5, 6})
	// 删除位于索引 i 的元素
	i = 3
	a = append(a[:i], a[i+1:]...)
	assert.Equal(t, a, []int{1, 2, 3, 5, 6})
	// 切除切片 a 中从索引 i 至 j 位置的元素 [i, j)
	a = []int{1, 2, 3, 4, 5, 6}
	i = 2
	j = 4
	a = append(a[:i], a[j:]...)
	assert.Equal(t, a, []int{1, 2, 5, 6})
	// 为切片 a 扩展 j 个元素长度
	a = []int{1, 2, 3, 4, 5, 6}
	j = 4
	a = append(a, make([]int, 4)...)
	assert.Equal(t, a, []int{1, 2, 3, 4, 5, 6, 0, 0, 0, 0})
	// 在索引 i 的位置插入元素 x
	x = 4
	i = 3
	a = []int{1, 2, 3, 5, 6}
	a = append(a[:i], append([]int{x}, a[i:]...)...)
	assert.Equal(t, a, []int{1, 2, 3, 4, 5, 6})
	// 在索引 i 的位置插入长度为 j 的新切片
	a = []int{1, 2, 5, 6}
	i = 2
	j = 2
	a = append(a[:i], append(make([]int, 2), a[i:]...)...)
	assert.Equal(t, a, []int{1, 2, 0, 0, 5, 6})
	// 在索引 i 的位置插入切片 b 的所有元素
	b = []int{3, 4}
	a = []int{1, 2, 5, 6}
	i = 2
	j = 2
	a = append(a[:i], append(b, a[i:]...)...)
	assert.Equal(t, a, []int{1, 2, 3, 4, 5, 6})
	// 取出位于切片 a 最末尾的元素 x
	a = []int{1, 2, 3, 4, 5, 6}
	x, a = a[len(a)-1], a[:len(a)-1]
	assert.Equal(t, x, 6)
	assert.Equal(t, a, []int{1, 2, 3, 4, 5})
	// 将元素 x 追加到切片 a
	a = append(a, 6)
	assert.Equal(t, a, []int{1, 2, 3, 4, 5, 6})
}
