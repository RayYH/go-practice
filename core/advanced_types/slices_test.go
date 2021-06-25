package advanced_types

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBasicUsagesOfSlices(t *testing.T) {
	// this declaration will produce an anonymous array
	b := []byte{'g', 'o', 'l', 'a', 'n', 'g'}
	assert.Equal(t, "[111 108 97]", fmt.Sprint(b[1:4]))
	assert.Equal(t, "[103 111]", fmt.Sprint(b[:2]))
	assert.Equal(t, "[108 97 110 103]", fmt.Sprint(b[2:]))
	assert.Equal(t, "[103 111 108 97 110 103]", fmt.Sprint(b[:]))
}

// make only makes slices, maps, and channels
// new only returns pointers to initialised memory
func TestSlicesInitialization(t *testing.T) {
	// literal
	names := []string{"leo", "jessica", "paul"}
	assert.Equal(t, len(names), 3)
	assert.Equal(t, cap(names), 3)

	// capacity=length=10
	checks := make([]bool, 10)
	assert.Equal(t, len(checks), 10)
	assert.Equal(t, cap(checks), 10)

	// capacity=2, length=20
	scores := make([]int, 2, 20)
	assert.Equal(t, len(scores), 2)
	assert.Equal(t, cap(scores), 20)

	// if we don't specified [0:2], numbers will be an array instead of a slice
	numbers := new([20]int)[0:2]
	assert.Equal(t, len(numbers), 2)
	assert.Equal(t, cap(numbers), 20)

	assert.Equal(t, "[leo jessica paul]", fmt.Sprint(names))
	assert.Equal(t, "[false false false false false false false false false false]", fmt.Sprint(checks))
	assert.Equal(t, "[0 0]", fmt.Sprint(scores))
	assert.Equal(t, "[0 0]", fmt.Sprint(numbers))
}

func TestSlicesBasicOperation(t *testing.T) {
	var arr [6]int
	assert.Equal(t, len(arr), 6)

	slice := arr[2:5]

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

func TestSliceCapacityAndLength(t *testing.T) {
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[:0]
	assert.Equal(t, cap(s), 6)
	assert.Equal(t, len(s), 0)

	s = s[:4]
	assert.Equal(t, cap(s), 6)
	assert.Equal(t, len(s), 4)

	s = s[2:]
	// notice here, first two values has been dropped
	assert.Equal(t, cap(s), 4)
	assert.Equal(t, len(s), 2)
}

func TestNilSlice(t *testing.T) {
	var s []int
	assert.Equal(t, cap(s), 0)
	assert.Equal(t, len(s), 0)
	assert.Nil(t, s)
}

func TestSliceMovement(t *testing.T) {
	slice := []int{1, 2, 3, 4, 5, 6}
	slice = slice[:]
	assert.Equal(t, len(slice), 6)
	assert.Equal(t, slice[0], 1)
	slice = slice[1:]
	assert.Equal(t, len(slice), 5)
	assert.Equal(t, slice[0], 2)
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
	// you cannot modify chars inside a string
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

func TestAppendToSlicesBasic(t *testing.T) {
	var s []int
	assert.Nil(t, s)
	s = append(s, 0)
	assert.Equal(t, s, []int{0})
	s = append(s, 1, 2, 3)
	assert.Equal(t, s, []int{0, 1, 2, 3})
}

func TestAppendSlices(t *testing.T) {
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

	// insert a new empty slice with j elements at index i
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

func TestResizeSlices(t *testing.T) {
	scores := make([]int, 5)
	assert.Equal(t, cap(scores), 5)
	// To a compiler, you're telling it to append a value to a slice that already holds 5 values.
	scores = append(scores, 9332)
	scores = append(scores, 9331)
	assert.Equal(t, scores, []int{0, 0, 0, 0, 0, 9332, 9331})

	lists := make([]int, 5)
	for i := 0; i < 25; i++ {
		lists = append(lists, i)
	}
	// 5 --> 10 --> 20 --> 40
	assert.Equal(t, cap(lists), 40)
}
