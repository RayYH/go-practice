package functions

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Go 并不支持可选参数、默认参数、泛型
// 在 Go 中，函数是一等公民 (可以将函数赋给一个变量)

func TestFunctionDeclaration(t *testing.T) {
	// Go 中函数的定义使用 `func` 关键字
	// 我们这里直接在函数内部定义一个函数，而不是在直接在包下定义
	multiply3Nums := func(a, b, c int) int {
		return a * b * c
	}

	assert.Equal(t, 60, multiply3Nums(2, 5, 6))
}

func TestReturnValues(t *testing.T) {
	t.Run("function can return multi values", func(t *testing.T) {
		// 不同于 Java、PHP，Go 中的函数可以直接返回多个返回值
		getX2AndX3 := func(n int) (int, int) {
			return n * 2, n * 3
		}

		x2, x3 := getX2AndX3(2)
		assert.Equal(t, 4, x2)
		assert.Equal(t, 6, x3)
	})

	t.Run("return values can be named", func(t *testing.T) {
		// Go 中的函数可以返回具名返回值，拥有名称的返回值可以直接使用，不需要再次声明
		// 即使函数使用了命名返回值，你依旧可以无视它而返回明确的值
		getNamedX2AndX3 := func(n int) (x2, x3 int) {
			x2, x3 = n*2, n*3
			// 不带任何变量的 `return` 语句会返回具名返回值对应的变量，等价于 `return x2, x3`
			return
		}

		x2, x3 := getNamedX2AndX3(2)
		assert.Equal(t, 4, x2)
		assert.Equal(t, 6, x3)
	})

	// Go 中函数的返回值类型可以是另一个函数，我们通常称这种函数为工厂函数
	t.Run("type of return value can be function", func(t *testing.T) {
		makeAddSuffix := func(suffix string) func(string) string {
			return func(name string) string {
				if !strings.HasSuffix(name, suffix) {
					name += suffix
				}
				return name
			}
		}

		addBmp := makeAddSuffix(".bmp")
		addJpeg := makeAddSuffix(".jpeg")
		assert.Equal(t, "file.bmp", addBmp("file"))
		assert.Equal(t, "file.jpeg", addJpeg("file"))
	})
}

func TestParameters(t *testing.T) {
	// Go 默认使用按值传递来传递参数，我们说的传递指针，实际上传递的是指针的拷贝，所以可以修改其指向的值
	t.Run("params can be modified inside func by passing references", func(t *testing.T) {
		multiply := func(a, b int, reply *int) {
			*reply = a * b
		}

		n := 0
		// & 符号用于获取变量的地址
		reply := &n
		multiply(3, 4, reply)
		assert.Equal(t, 12, *reply)
	})

	// 同其他语言一样，Go 使用 ... 来表示剩余参数 (变长参数)
	t.Run("rest parameters", func(t *testing.T) {
		min := func(s ...int) int {
			if len(s) == 0 {
				return 0
			}

			minValue := s[0]

			for _, v := range s {
				if v < minValue {
					minValue = v
				}
			}

			return minValue
		}

		assert.Equal(t, 1, min(1, 2, 3, 4, 5))
		assert.Equal(t, 1, min([]int{1, 2, 3, 4, 5}...))
		assert.Equal(t, 2, min([]int{1, 2, 3, 4, 5}[1:]...))
		assert.Equal(t, 0, min())
		assert.Equal(t, -5, min(1, 2, -5, 3, 111))
	})

	// Go 中的空接口可以接受任意类型的变量作为参数
	t.Run("empty interface can accept param of any type", func(t *testing.T) {
		getType := func(arg interface{}) string {
			switch arg.(type) {
			case int:
				return "int"
			case string:
				return "string"
			case int64:
				return "int64"
			default:
				return "unknown"
			}
		}

		var v1 = 1 // inferred int
		var v2 int64 = 234
		var v3 = "Hello" // inferred string
		var v4 float32 = 1.234
		assert.Equal(t, "int", getType(v1))
		assert.Equal(t, "int64", getType(v2))
		assert.Equal(t, "string", getType(v3))
		assert.Equal(t, "unknown", getType(v4))
	})
}

func TestRecursion(t *testing.T) {
	result := 0
	nums := []int{1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89}
	for i := 0; i <= 10; i++ {
		result = fibonacci(i)
		assert.Equal(t, result, nums[i])
	}
}

func TestCallbacksAndClosures(t *testing.T) {
	// 作为参数传递的函数我们称之为回调
	// 闭包是一个可调用对象 (这个对象的表现形式通常就是函数)，它记录了一些信息，这些信息来自于创建它的作用域
	// 可以简单理解为使用了上下文变量的回调就是闭包
	t.Run("callback", func(t *testing.T) {
		add := func(a, b int) int {
			return a + b
		}

		add2 := func(a, b int) int {
			return a + b + a + b
		}

		addWrapper := func(x, y int, f func(int, int) int) int {
			return f(x, y)
		}

		// 这里 add 和 add2 作为参数传给 addWrapper，所以 add 和 add2 是回调函数
		assert.Equal(t, 3, addWrapper(1, 2, add))
		assert.Equal(t, 6, addWrapper(1, 2, add2))
	})

	t.Run("closure", func(t *testing.T) {
		multiplier := 2
		multiplyClosure := func(i int) int {
			// 闭包定义可以直接使用外部变量，注意是变量，而不是变量的值
			return i * multiplier
		}

		assert.Equal(t, 4, multiplyClosure(2))

		// 如果更改了变量的值，相应的闭包无需重新声明，会自动利用变量最新的值
		multiplier = 3
		assert.Equal(t, 6, multiplyClosure(2))
	})

	t.Run("anonymous functions", func(t *testing.T) {
		addWrapper := func(x, y int, f func(int, int) int) int {
			return f(x, y)
		}

		c := 100

		// 匿名函数同样被称之为闭包，它们被允许调用定义在其它环境下的变量
		// 闭包可使得某个函数捕捉到一些外部状态 (一个闭包继承了函数所声明时的作用域)
		assert.Equal(t, 105, addWrapper(2, 3, func(x, y int) int {
			return x + y + c
		}))
	})

	t.Run("adder example", func(t *testing.T) {
		addTwo := func() func(int) int {
			return func(b int) int {
				return b + 2
			}
		}

		adder := func(a int) func(int) int {
			return func(b int) int {
				return a + b
			}
		}

		plusTwo := addTwo()
		assert.Equal(t, 5, plusTwo(3))
		plus := adder(4)
		assert.Equal(t, 9, plus(5))
	})

	t.Run("holding reference", func(t *testing.T) {
		sequentialAdder := func() func(int) int {
			var x int
			// 在多次调用中，变量 x 的值是被保留的
			// 闭包函数保存并积累其中的变量的值，不管外部函数退出与否，它都能够继续操作外部函数中的局部变量
			return func(i int) int {
				// 闭包中的变量可以理解为持有外部变量的引用，因此对该函数调用多次，值依然是一样的
				x += i
				return x
			}
		}

		var f = sequentialAdder()
		assert.Equal(t, 1, f(1))
		assert.Equal(t, 21, f(20))
		assert.Equal(t, 321, f(300))
	})
}
