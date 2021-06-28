package advanced_types

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// 数组遍历
func TestArraysIteration(t *testing.T) {
	// 数组声明之后，数组中的每个元素都具有对应类型的零值
	// 数组长度也是数组类型的一部分，所以 [5]int 和 [10]int 属于不同类型
	var intArr [5]int

	// 我们可以使用 C 风格的循环体来操作数组，len(intArr) 返回数组 intArr 的长度
	for i := 0; i < len(intArr); i++ {
		// 通过索引下标我们可以访问和修改数组中的元素
		assert.Equal(t, intArr[i], 0)
		intArr[i] = i * 2
	}

	// for range 风格可以同时获取元素的 k/v，如果指向获取 k 或者 v，使用空白标识符 _
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

// 数组字面量
func TestArraysLiterals(t *testing.T) {
	t.Run("declaration and initialization", func(t *testing.T) {
		// 声明并同时进行初始化
		var arr1 = [5]int{18, 20, 15, 22, 16}
		// 先声明，随后再进行初始化
		var arr2 [5]int
		arr2 = [5]int{18, 20, 15, 22, 16}
		assert.Equal(t, arr1, arr2)
	})

	t.Run("[...] and []", func(t *testing.T) {
		// 不指定具体长度，使用 ... 时，表示由 {} 字面量里的元素来反映数组的真实长度
		// 如果使用 var emptyArr [...]int，则会报错：use of [...] array outside of array literal
		// 也就是说，[...] 只使用同时初始化的情况 (只有初始化才能明确具体的数组长度)
		var arr = [...]int{5, 6, 7, 8, 22} // array
		assert.Equal(t, "[5]int", fmt.Sprint(reflect.TypeOf(arr)))
		// 不显式指出长度则为切片
		var sli = []int{5, 6, 7, 8, 22} // slice
		assert.Equal(t, "[]int", fmt.Sprint(reflect.TypeOf(sli)))
		assert.Equal(t, cap(sli), 5)
	})

	t.Run("skip indexes", func(t *testing.T) {
		// 只有下标 3 和 4 赋予了元素，因此数组和切片的长度都是 5，下标 0 1 2 都为零值 0
		var arr = [5]string{3: "Chris", 4: "Ron"} // array
		var sli = []string{3: "Chris", 4: "Ron"}  // slice

		// 总结：[count] 和 [...] 代表数组，[] 代表切片
		assert.Equal(t, 5, cap(arr))
		assert.Equal(t, 5, len(arr))
		assert.Equal(t, 5, cap(sli))
		assert.Equal(t, 5, len(sli))
	})

	// 如果我们想让数组元素类型为任意类型的话可以使用空接口作为类型
	// 当使用值时我们必须先做一个类型判断
	t.Run("empty interface", func(t *testing.T) {
		var anyType = [...]interface{}{"1", 2, true}
		assert.Equal(t, cap(anyType), 3)
	})
}

// cap 方法，返回数组或者切片对应的数组的长度，详见 slice 相关示例
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
		// 接收 [3]int 数组作为参数
		funcAcceptArray := func(arr [3]int) string {
			arr[0] = 1
			return fmt.Sprint(arr)
		}

		// 接收指向 [3]int 数组的指针作为参数
		funcAcceptArrayPointer := func(arr *[3]int) string {
			arr[0] = 2
			return fmt.Sprint(arr)
		}

		var arr [3]int

		// arr 没有改变，因为传递的是数组的拷贝
		assert.Equal(t, "[1 0 0]", funcAcceptArray(arr))
		assert.Equal(t, [3]int{0, 0, 0}, arr)

		// 值发生了改变，因为传递的是指针的拷贝
		assert.Equal(t, "&[2 0 0]", funcAcceptArrayPointer(&arr))
		assert.Equal(t, [3]int{2, 0, 0}, arr)
	})
}

// Go 中所有的参数都是按值传递的，为了减少值拷贝的开销，在接收数组作为参数时
// 我们可以将数组指针作为参数传递，除此之外，我们可以传递切片，切片本质上是一个特殊的结构体
// 该结构体里面包含了一个指向底层数组的指针，因此传递的值的大小与数组大小无关
// type SliceHeader struct {
//	Data uintptr
//	Len  int
//	Cap  int
//}
func TestArraysPointerOrSlicesAsArguments(t *testing.T) {
	// 传递数组的指针
	Sum := func(numbers *[3]float64) (sum float64) {
		for _, v := range numbers {
			sum += v
		}

		return
	}

	// 传递切片
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
