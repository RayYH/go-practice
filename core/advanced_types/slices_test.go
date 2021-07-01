package advanced_types

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// 切片是对数组一个连续片段的引用，这里的数组我们称之为关联数组，通常是匿名的
// 切片本质上是一个特殊的结构体，该结构体里面包含了一个指向底层数组的指针，因此传递的值的大小与数组大小无关
// type SliceHeader struct {
//	Data uintptr
//	Len  int
//	Cap  int
//}

func TestSlicesInitialization(t *testing.T) {
	// 可以直接使用字面量声明一个切片：v := []type{...}
	t.Run("using literals", func(t *testing.T) {
		names := []string{"leo", "jessica", "paul"}
		assert.Equal(t, len(names), 3)
		assert.Equal(t, cap(names), 3)

		assert.Equal(t, "[leo jessica paul]", fmt.Sprint(names))
	})

	// Go 中 make 只能用于创建切片、map、通道 (初始化为零值)，而 new 返回一个已经初始化内存的指针
	// new 适用于值类型，比如数组和结构体
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
		// 这里的 [0:2] 表示为数组 [20]int 的一个切片
		numbers := new([20]int)[0:2]
		assert.Equal(t, len(numbers), 2)
		assert.Equal(t, cap(numbers), 20)

		assert.Equal(t, "[0 0]", fmt.Sprint(numbers))
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

	// 你可以像操作数组那样操作切片
	t.Run("indexes of slice", func(t *testing.T) {
		var arr [6]int
		slice := arr[2:5]

		// 赋值
		for i := 0; i < len(arr); i++ {
			arr[i] = i * 2
		}
		// 访问
		for i := 0; i < len(slice); i++ {
			// 注意切片中的索引值与关联数组的索引值并不一致
			// slice[0,1,2,3] -> arr[2,3,4,5]
			assert.Equal(t, slice[i], (i+2)*2)
		}
	})

	// 如果多个切片引用的是同一个数组，则它们之间 (包括数组本身) 是共享数据的
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

	// 切片的长度永远不会超过它的容量，所以对于切片 s 来说 0 <= len(s) <= cap(s) 永远成立
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

	// 一个切片在未初始化之前默认为 nil，长度为 0
	t.Run("empty slice", func(t *testing.T) {
		var s []int
		assert.Equal(t, cap(s), 0)
		assert.Equal(t, len(s), 0)
		assert.Nil(t, s)
	})
}

func TestOperations(t *testing.T) {
	t.Run("movement", func(t *testing.T) {
		// slice 是匿名数组 []int{1, 2, 3, 4, 5, 6} 的一个完整切片
		slice := []int{1, 2, 3, 4, 5, 6}
		assert.Equal(t, len(slice), 6)
		assert.Equal(t, slice[0], 1)
		// 这里基于 slice 继续切片 - 将 slice 向后移动一位
		slice = slice[1:]
		assert.Equal(t, len(slice), 5)
		assert.Equal(t, slice[0], 2)
	})

	// 改变切片的过程称之为切片重组 (reslice)
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

		// 我们可以在声明切片的过程中显式指出容量
		t.Run("three index", func(t *testing.T) {
			s := ar[2:4:7]
			assert.Equal(t, 2, len(s))
			assert.Equal(t, 5, cap(a))
		})
	})

	// 通过 copy 方法来对切片进行复制
	t.Run("duplication", func(t *testing.T) {
		// use copy
		slFrom := []int{1, 2, 3}
		slTo := make([]int, 10)
		n := copy(slTo, slFrom) // copy func returned copied elements
		assert.Equal(t, 3, n)
		assert.Equal(t, slTo, []int{1, 2, 3, 0, 0, 0, 0, 0, 0, 0})
		// 复制之后的两个切片没有任何关联
		slFrom[2] = 9
		assert.Equal(t, slTo[2], 3)
	})
}

func TestSliceAndStrings(t *testing.T) {
	t.Run("generate slices from strings", func(t *testing.T) {
		s := "\u00ff\u754c"
		// []byte(s) 将字符串转换为字符切片，处理 UTF-8 字符时使用 []int32 或者 []rune
		// 此时可以通过代码 len([]int32(s)) 来获得字符串中字符的数量，但使用 utf8.RuneCountInString(s) 效率会更高一点
		b := []byte(s)
		res := ""
		for i, c := range s {
			assert.Equal(t, b[i], s[i])
			res += fmt.Sprintf("%d:%c ", i, c)
		}

		assert.Equal(t, strings.TrimSpace(res), "0:ÿ 2:界")
	})

	// 字符串中的字符是不可变的，为了达到这个目的，我们需要将字符串转换为字符数组并进行修改，然后再转换为字符串
	t.Run("strings are immutable but byte slices are mutable", func(t *testing.T) {
		s := "Hello"
		c := []byte(s)
		c[0] = 'c'
		assert.Equal(t, string(c), "cello")
	})

	// 示例：字节切片 (数组) 的比较函数
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

	// append 最基本的用法就是往切片中追加元素，你不必考虑容量不够的问题，切片会自动扩容
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

	// 用 ... 解构运算符配合 append 可以将切片中的所有元素追加到另一个切片
	t.Run("append all elements", func(t *testing.T) {
		a := []int{1, 2, 3}
		b := []int{4, 5, 6}
		a = append(a, b...)
		assert.Equal(t, a, []int{1, 2, 3, 4, 5, 6})
		// 复制后的元素与原切片不再有任何联系
		b[0] = 9
		assert.Equal(t, a, []int{1, 2, 3, 4, 5, 6})
	})

	// 从切片中删除指定索引处的元素
	t.Run("delete element", func(t *testing.T) {
		a := []int{1, 2, 3, 4, 5, 6}
		assert.Equal(t, cap(a), 6)
		i = 3
		a = append(a[:i], a[i+1:]...)
		assert.Equal(t, a, []int{1, 2, 3, 5, 6})
		// 注意到切片的容量并没有发生变化
		assert.Equal(t, cap(a), 6)
	})

	// 从切片中删除指定范围内的元素
	t.Run("delete elements of specified range", func(t *testing.T) {
		a := []int{1, 2, 3, 4, 5, 6}
		i, j = 2, 4 // remove elements of index 2, 3
		a = append(a[:i], a[j:]...)
		assert.Equal(t, a, []int{1, 2, 5, 6})
	})

	// 在切片末尾追加零值元素
	t.Run("extends elements", func(t *testing.T) {
		a := []int{1, 2, 3, 4, 5, 6}
		j = 4
		a = append(a, make([]int, 4)...)
		assert.Equal(t, a, []int{1, 2, 3, 4, 5, 6, 0, 0, 0, 0})
	})

	// 在切片指定位置插入元素
	t.Run("insert element at specified position", func(t *testing.T) {
		e = 4
		i = 3
		a := []int{1, 2, 3, 5, 6}
		a = append(a[:i], append([]int{e}, a[i:]...)...)
		assert.Equal(t, a, []int{1, 2, 3, 4, 5, 6})
	})

	// 在指定位置插入多个元素
	t.Run("insert a new slice with zero elements at specified position", func(t *testing.T) {
		a := []int{1, 2, 5, 6}
		i = 2
		j = 3
		a = append(a[:i], append(make([]int, j), a[i:]...)...)
		assert.Equal(t, a, []int{1, 2, 0, 0, 0, 5, 6})
	})

	// 在指定位置插入另一个切片的所有元素
	t.Run("insert all elements from one slice to another at specified position", func(t *testing.T) {
		a := []int{1, 2, 5, 6}
		b := []int{3, 4}
		i = 2
		a = append(a[:i], append(b, a[i:]...)...)
		assert.Equal(t, a, []int{1, 2, 3, 4, 5, 6})

	})

	// 获取并弹出切片的最后一个元素
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

	// 新的长度不超过原长度的两倍，直接按愿长度的两倍扩容
	t.Run("the new slice is less than twice the old slice capacity", func(t *testing.T) {
		scores := make([]int, 5)
		// 新增一个元素，容量扩大为两倍 5 -> 10
		scores = append(scores, 1)
		assert.Equal(t, cap(scores), 10)

		scores2 := make([]int, 9)
		scores2 = append(scores2, 1)
		assert.Equal(t, cap(scores2), 18)

		// TODO: 为啥是 40 而不是 38
		scores3 := make([]int, 19)
		scores3 = append(scores3, 1)
		assert.Equal(t, cap(scores3), 40)

		// TODO: 为啥是 1696 而不是 1554
		scores4 := make([]int, 777)
		scores4 = append(scores4, 1)
		assert.Equal(t, cap(scores4), 1696)
	})

	// 新的长度超过原长度的两倍，采用新的长度
	// 注意到，这里扩容与 CPU 有关，64 位 CPU 中可以保存两个 int32 类型 (数据结构对齐)
	// 因此下面 4 个 int32 增加 5 个 int32 后，长度是 10 而不是 9
	t.Run("the capacity after multiplying by 2 is still less than the new slice capacity", func(t *testing.T) {
		if is64Bit {
			scores := make([]int, 4)
			assert.Equal(t, cap(scores), 4)
			scores = append(scores, make([]int, 5)...)
			assert.Equal(t, cap(scores), 10)

			bytes := make([]byte, 8) // 8*8=64
			assert.Equal(t, cap(bytes), 8)
			bytes = append(bytes, make([]byte, 9)...) // 8*3=24 > 8+9=17
			// TODO: 不同版本实现不同
			// 1.16 -> 24
			// 1.13-1.15 -> 32
			assert.GreaterOrEqual(t, cap(bytes), 24)
		}
	})

	// 超过 1024 按 1.25 倍扩容
	t.Run("old slice capacity is greater than 1024", func(t *testing.T) {
		scores := make([]int, 1024)
		assert.Equal(t, 1024, cap(scores))
		scores = append(scores, make([]int, 1)...)
		// 1024*1.25 = 1280
		assert.Equal(t, cap(scores), 1280)
	})

	// 如果切片的容量不足以存储新增元素，append 会分配新的切片来保证已有切片元素和新增元素的存储
	// 因此，返回的切片可能已经指向一个不同的相关数组了
	t.Run("share same data if slices have no need to grow", func(t *testing.T) {
		a := []int{1, 2}
		// 切片扩容，2 --> 4
		b := append(a, 3)

		// 由于 c 和 d 都不会触发扩容，因此它们共享同一个底层数组，所以 c[3] 和 d[3] 都是 5
		c := append(b, 4)
		d := append(b, 5)
		assert.Equal(t, 5, c[3])
		assert.Equal(t, 5, d[3])

		b[0] = 3
		assert.Equal(t, 3, c[0])
		assert.Equal(t, 3, d[0])
	})
}
