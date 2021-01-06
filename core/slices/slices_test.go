package slices

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestBasicUsagesOfSlices(t *testing.T) {
	// will produce an anonymous array for this slice
	b := []byte{'g', 'o', 'l', 'a', 'n', 'g'}
	assert.Equal(t, "[111 108 97]", fmt.Sprint(b[1:4]))
	assert.Equal(t, "[103 111]", fmt.Sprint(b[:2]))
	assert.Equal(t, "[108 97 110 103]", fmt.Sprint(b[2:]))
	assert.Equal(t, "[103 111 108 97 110 103]", fmt.Sprint(b[:]))
}

// make only makes slices, maps, and channels
// new only returns pointers to initialised memory
func TestSlicesInitialization(t *testing.T) {
	names := []string{"leo", "jessica", "paul"}
	checks := make([]bool, 10)
	scores := make([]int, 2, 20)
	numbers := new([20]int)[0:2]
	assert.Equal(t, "[leo jessica paul]", fmt.Sprint(names))
	assert.Equal(t, "[false false false false false false false false false false]", fmt.Sprint(checks))
	assert.Equal(t, "[0 0]", fmt.Sprint(scores))
	assert.Equal(t, "[0 0]", fmt.Sprint(numbers))
}

func TestSliceUsage(t *testing.T) {
	var arr [6]int
	assert.Equal(t, len(arr), 6)

	// slice from arr
	var slice = arr[2:5]

	for i := 0; i < len(arr); i++ {
		arr[i] = i * 2
	}

	for i := 0; i < len(slice); i++ {
		assert.Equal(t, slice[i], (i+2)*2)
	}

	assert.Equal(t, len(slice), 3) // [2:5] --> 2, 3, 4
	assert.Equal(t, cap(slice), 4) // arr[0:6] --> 2, 3, 4, 5
}

func TestSlicesShareCommonData(t *testing.T) {
	// if multi slices refs same array
	// the array elements are shared between them (including array self)
	// actually, array is the base block which slices generated from
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
	SliceForRange()
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
	ReSlice()
	// Output:
	// The length of slice[0:1] is 1
	// The length of slice[0:2] is 2
	// The length of slice[0:3] is 3
	// The length of slice[0:4] is 4
	// The length of slice[0:5] is 5
	// The length of slice[0:6] is 6
	// The length of slice[0:7] is 7
	// The length of slice[0:8] is 8
	// The length of slice[0:9] is 9
	// The length of slice[0:10] is 10
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

	var a = ar[5:7] // ref to subarray {5,6} - len(a) is 2 and cap(a) is 5
	assert.Equal(t, a, []int{5, 6})
	assert.Equal(t, 2, len(a))
	assert.Equal(t, 5, cap(a))

	a = a[0:4] // ref to subarray {5,6,7,8} - len(a) is now 4 but cap(a) is still 5
	assert.Equal(t, a, []int{5, 6, 7, 8})
	assert.Equal(t, 4, len(a))
	assert.Equal(t, 5, cap(a))
}

func TestSlicesDuplication(t *testing.T) {
	// use copy
	slFrom := []int{1, 2, 3}
	slTo := make([]int, 10)
	n := copy(slTo, slFrom) // copy func returned copied elements
	assert.Equal(t, 3, n)
	assert.Equal(t, slTo, []int{1, 2, 3, 0, 0, 0, 0, 0, 0, 0})
	// use append
	slice := []int{1, 2, 3}
	slice = append(slice, 4, 5, 6)
	assert.Equal(t, slice, []int{1, 2, 3, 4, 5, 6})
}

func TestGenerateSlicesFromStrings(t *testing.T) {
	s := "\u00ff\u754c"
	res := ""
	for i, c := range s {
		res += fmt.Sprintf("%d:%c ", i, c)
	}
	assert.Equal(t, strings.TrimSpace(res), "0:ÿ 2:界")
}

func TestAppendStringsToByteArray(t *testing.T) {
	var b []byte
	b = []byte{'H', 'e', 'l', 'l', 'o'}
	var s string
	s = " World"
	b = append(b, s...)
	assert.Equal(t, string(b), "Hello World")
	assert.Equal(t, b, []byte{'H', 'e', 'l', 'l', 'o', ' ', 'W', 'o', 'r', 'l', 'd'})
}

func TestStringsAreImmutable(t *testing.T) {
	s := "Hello"
	c := []byte(s)
	c[0] = 'c'
	assert.Equal(t, string(c), "cello")
}

func TestCompareSlices(t *testing.T) {
	compare := func(a, b []byte) int {
		for i := 0; i < len(a) && i < len(b); i++ {
			switch {
			case a[i] > b[i]:
				return 1
			case a[i] < b[i]:
				return -1
			}
		}
		// check length
		switch {
		case len(a) < len(b):
			return -1
		case len(a) > len(b):
			return 1
		}
		return 0
	}
	assert.Equal(t, compare([]byte{1, 2, 3}, []byte{1, 2, 3}), 0)
	assert.Equal(t, compare([]byte{1, 2, 4}, []byte{1, 2, 3}), 1)
	assert.Equal(t, compare([]byte{1, 2, 3}, []byte{1, 2, 4}), -1)
	assert.Equal(t, compare([]byte{1, 2, 3, 4}, []byte{1, 2, 3}), 1)
	assert.Equal(t, compare([]byte{1, 2, 3}, []byte{1, 2, 3, 4}), -1)
}

func TestAppend(t *testing.T) {
	var i, j int
	var x int
	a := []int{1, 2, 3}
	b := []int{4, 5, 6}
	// append all elements of b
	a = append(a, b...)
	assert.Equal(t, a, []int{1, 2, 3, 4, 5, 6})

	// create a new slice b from existed slice a
	b = make([]int, len(a))
	copy(b, a)
	assert.Equal(t, b, []int{1, 2, 3, 4, 5, 6})

	// delete element at index i
	i = 3
	a = append(a[:i], a[i+1:]...)
	assert.Equal(t, a, []int{1, 2, 3, 5, 6})

	// remove a[i, j)
	a = []int{1, 2, 3, 4, 5, 6}
	i = 2
	j = 4
	a = append(a[:i], a[j:]...)
	assert.Equal(t, a, []int{1, 2, 5, 6})

	// extends j elements for slice a
	a = []int{1, 2, 3, 4, 5, 6}
	j = 4
	a = append(a, make([]int, 4)...)
	assert.Equal(t, a, []int{1, 2, 3, 4, 5, 6, 0, 0, 0, 0})

	// insert element x at index i
	x = 4
	i = 3
	a = []int{1, 2, 3, 5, 6}
	a = append(a[:i], append([]int{x}, a[i:]...)...)
	assert.Equal(t, a, []int{1, 2, 3, 4, 5, 6})

	// insert a new empty slice at index i
	a = []int{1, 2, 5, 6}
	i = 2
	j = 3
	a = append(a[:i], append(make([]int, j), a[i:]...)...)
	assert.Equal(t, a, []int{1, 2, 0, 0, 0, 5, 6})

	// insert all elements from b at index i of a
	b = []int{3, 4}
	a = []int{1, 2, 5, 6}
	i = 2
	a = append(a[:i], append(b, a[i:]...)...)
	assert.Equal(t, a, []int{1, 2, 3, 4, 5, 6})

	// fetch last element x of slice a
	a = []int{1, 2, 3, 4, 5, 6}
	x, a = a[len(a)-1], a[:len(a)-1]
	assert.Equal(t, x, 6)
	assert.Equal(t, a, []int{1, 2, 3, 4, 5})

	// append element x to slice a
	a = append(a, 6)
	assert.Equal(t, a, []int{1, 2, 3, 4, 5, 6})
}
