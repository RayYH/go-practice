package advanced_types

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSlicesInitialization(t *testing.T) {
	t.Run("using literals", func(t *testing.T) {
		names := []string{"leo", "jessica", "paul"}
		assert.Equal(t, len(names), 3)
		assert.Equal(t, cap(names), 3)

		assert.Equal(t, "[leo jessica paul]", fmt.Sprint(names))
	})

	t.Run("using make method", func(t *testing.T) {
		// capacity=length=10
		checks := make([]bool, 10)
		assert.Equal(t, len(checks), 10)
		assert.Equal(t, cap(checks), 10)
		assert.Equal(t, "[false false false false false false false false false false]", fmt.Sprint(checks))

		// capacity=20, length=2
		scores := make([]int, 2, 20)
		assert.Equal(t, len(scores), 2)
		assert.Equal(t, cap(scores), 20)

		assert.Equal(t, "[0 0]", fmt.Sprint(scores))
	})

	t.Run("using new method", func(t *testing.T) {
		numbers := new([20]int)[0:2]
		assert.Equal(t, len(numbers), 2)
		assert.Equal(t, cap(numbers), 20)

		assert.Equal(t, "[0 0]", fmt.Sprint(numbers))
	})

	t.Run("three index", func(t *testing.T) {
		var ar = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
		s := ar[2:4:7]
		assert.Equal(t, 2, len(s))
		assert.Equal(t, 5, cap(s))
	})
}

func TestPropertiesAndFeatures(t *testing.T) {
	t.Run("range of slice", func(t *testing.T) {
		b := []byte{'g', 'o', 'l', 'a', 'n', 'g'}
		assert.Equal(t, "[111 108 97]", fmt.Sprint(b[1:4]))           // indexes of 1 2 3
		assert.Equal(t, "[103 111]", fmt.Sprint(b[:2]))               // indexes of 0 1
		assert.Equal(t, "[108 97 110 103]", fmt.Sprint(b[2:]))        // indexes of 2 3 4 5
		assert.Equal(t, "[103 111 108 97 110 103]", fmt.Sprint(b[:])) // indexes of 0 1 2 3 4 5
	})

	t.Run("indexes of slice", func(t *testing.T) {
		var arr [6]int
		slice := arr[2:5]

		// assignment
		for i := 0; i < len(arr); i++ {
			arr[i] = i * 2
		}
		// access
		for i := 0; i < len(slice); i++ {
			// slice[0,1,2,3] -> arr[2,3,4,5]
			assert.Equal(t, slice[i], (i+2)*2)
		}
	})

	t.Run("slices refer to same array share common data", func(t *testing.T) {
		values := []int{1, 2, 3, 4, 5, 6}
		slice1 := values[1:]
		slice2 := values[1:]
		assert.Equal(t, slice1[0], 2)
		assert.Equal(t, slice2[0], 2)
		values[1] = 9
		assert.Equal(t, values[1], 9)
		assert.Equal(t, slice1[0], 9)
		assert.Equal(t, slice2[0], 9)
	})

	t.Run("cap and len method", func(t *testing.T) {
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
	})

	t.Run("empty slice", func(t *testing.T) {
		var s []int
		assert.Equal(t, cap(s), 0)
		assert.Equal(t, len(s), 0)
		assert.Nil(t, s)
	})
}

func TestOperations(t *testing.T) {
	t.Run("movement", func(t *testing.T) {
		slice := []int{1, 2, 3, 4, 5, 6}
		assert.Equal(t, len(slice), 6)
		assert.Equal(t, slice[0], 1)
		slice = slice[1:]
		assert.Equal(t, len(slice), 5)
		assert.Equal(t, slice[0], 2)
	})

	t.Run("reslice", func(t *testing.T) {
		var ar = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

		var a = ar[5:7] // ref to subarray {5,6} - len(a) is 2 and cap(a) is 5
		assert.Equal(t, a, []int{5, 6})
		assert.Equal(t, 2, len(a))
		assert.Equal(t, 5, cap(a))

		a = a[0:4] // ref to subarray {5,6,7,8} - len(a) is now 4 but cap(a) is still 5
		assert.Equal(t, a, []int{5, 6, 7, 8})
		assert.Equal(t, 4, len(a))
		assert.Equal(t, 5, cap(a))
	})

	t.Run("duplication", func(t *testing.T) {
		slFrom := []int{1, 2, 3}
		slTo := make([]int, 10)
		n := copy(slTo, slFrom) // copy func returned copied elements
		assert.Equal(t, 3, n)
		assert.Equal(t, slTo, []int{1, 2, 3, 0, 0, 0, 0, 0, 0, 0})
		slFrom[2] = 9
		assert.Equal(t, slTo[2], 3)
	})
}

func TestSliceAndStrings(t *testing.T) {
	t.Run("generate slices from strings", func(t *testing.T) {
		s := "\u00ff\u754c"
		b := []byte(s)
		assert.NotNil(t, b)
	})

	t.Run("strings are immutable but byte slices are mutable", func(t *testing.T) {
		s := "Hello"
		c := []byte(s)
		c[0] = 'c'
		assert.Equal(t, string(c), "cello")
	})

	t.Run("compare byte slices", func(t *testing.T) {
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
	})
}

func TestAppendSlices(t *testing.T) {
	var i, j, e int

	t.Run("append element(s)", func(t *testing.T) {
		var s []int
		assert.Nil(t, s)
		assert.Equal(t, cap(s), 0)
		s = append(s, 0)
		assert.Equal(t, s, []int{0})
		assert.Equal(t, cap(s), 1)
		s = append(s, 1, 2, 3)
		assert.Equal(t, s, []int{0, 1, 2, 3})
		assert.Equal(t, cap(s), 4)
	})

	t.Run("append all elements", func(t *testing.T) {
		a := []int{1, 2, 3}
		b := []int{4, 5, 6}
		a = append(a, b...)
		assert.Equal(t, a, []int{1, 2, 3, 4, 5, 6})
		b[0] = 9
		assert.Equal(t, a, []int{1, 2, 3, 4, 5, 6})
	})

	t.Run("delete element", func(t *testing.T) {
		a := []int{1, 2, 3, 4, 5, 6}
		assert.Equal(t, cap(a), 6)
		i = 3
		a = append(a[:i], a[i+1:]...)
		assert.Equal(t, a, []int{1, 2, 3, 5, 6})
		assert.Equal(t, cap(a), 6)
	})

	t.Run("delete elements of specified range", func(t *testing.T) {
		a := []int{1, 2, 3, 4, 5, 6}
		i, j = 2, 4 // remove elements of index 2, 3
		a = append(a[:i], a[j:]...)
		assert.Equal(t, a, []int{1, 2, 5, 6})
	})

	t.Run("extends elements", func(t *testing.T) {
		a := []int{1, 2, 3, 4, 5, 6}
		j = 4
		a = append(a, make([]int, 4)...)
		assert.Equal(t, a, []int{1, 2, 3, 4, 5, 6, 0, 0, 0, 0})
	})

	t.Run("insert element at specified position", func(t *testing.T) {
		e = 4
		i = 3
		a := []int{1, 2, 3, 5, 6}
		a = append(a[:i], append([]int{e}, a[i:]...)...)
		assert.Equal(t, a, []int{1, 2, 3, 4, 5, 6})
	})

	t.Run("insert a new slice with zero elements at specified position", func(t *testing.T) {
		a := []int{1, 2, 5, 6}
		i = 2
		j = 3
		a = append(a[:i], append(make([]int, j), a[i:]...)...)
		assert.Equal(t, a, []int{1, 2, 0, 0, 0, 5, 6})
	})

	t.Run("insert all elements from one slice to another at specified position", func(t *testing.T) {
		a := []int{1, 2, 5, 6}
		b := []int{3, 4}
		i = 2
		a = append(a[:i], append(b, a[i:]...)...)
		assert.Equal(t, a, []int{1, 2, 3, 4, 5, 6})

	})

	t.Run("pop last element of given slice", func(t *testing.T) {
		a := []int{1, 2, 3, 4, 5, 6}
		e, a = a[len(a)-1], a[:len(a)-1]
		assert.Equal(t, e, 6)
		assert.Equal(t, a, []int{1, 2, 3, 4, 5})
	})
}

// 切片容量增长与实现有关，网上的相关资料也较少，这里暂不深入研究
func TestCapacityGrows(t *testing.T) {
	const is64Bit = uint64(^uintptr(0)) == ^uint64(0)

	t.Run("the new slice is less than twice the old slice capacity", func(t *testing.T) {
		scores := make([]int, 5)
		scores = append(scores, 1)
		assert.Equal(t, cap(scores), 10)

		scores2 := make([]int, 9)
		scores2 = append(scores2, 1)
		assert.Equal(t, cap(scores2), 18)

		// TODO: why 40 not 38
		scores3 := make([]int, 19)
		scores3 = append(scores3, 1)
		assert.Equal(t, cap(scores3), 40)

		// TODO: why 1696 not 1554
		scores4 := make([]int, 777)
		scores4 = append(scores4, 1)
		assert.Equal(t, cap(scores4), 1696)
	})

	t.Run("the capacity after multiplying by 2 is still less than the new slice capacity", func(t *testing.T) {
		if is64Bit {
			scores := make([]int, 4)
			assert.Equal(t, cap(scores), 4)
			scores = append(scores, make([]int, 5)...)
			assert.Equal(t, cap(scores), 10)

			bytes := make([]byte, 8) // 8*8=64
			assert.Equal(t, cap(bytes), 8)
			bytes = append(bytes, make([]byte, 9)...) // 8*3=24 > 8+9=17
			// version 1.16 -> 24
			// version 1.13-1.15 -> 32
			assert.GreaterOrEqual(t, cap(bytes), 24)
		}
	})

	t.Run("old slice capacity is greater than 1024", func(t *testing.T) {
		scores := make([]int, 1024)
		assert.Equal(t, 1024, cap(scores))
		scores = append(scores, make([]int, 1)...)
		// 1024*1.25 = 1280
		assert.Equal(t, cap(scores), 1280)
	})

	t.Run("share same data if slices have no need to grow", func(t *testing.T) {
		a := []int{1, 2}
		b := append(a, 3)
		c := append(b, 4)
		d := append(b, 5)
		assert.Equal(t, 5, c[3])
		assert.Equal(t, 5, d[3])

		b[0] = 3
		assert.Equal(t, 3, c[0])
		assert.Equal(t, 3, d[0])
	})
}
