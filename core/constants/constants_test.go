package constants

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// 常量使用关键字 const 定义，用于存储不会改变的数据
func TestDeclarations(t *testing.T) {
	// 可以在函数内部声明常量，一旦函数退出，常量就会被释放
	t.Run("define constants via const keyword", func(t *testing.T) {
		// 常量的类型只能是布尔型、数值型 (整数、浮点数、复数)、字符串型
		const Pi = 3.14159
		assert.Equal(t, 3.14159, Pi)
	})

	// 可以在一行同时定义多个常量 (并行赋值)
	t.Run("constants can be declared at one line", func(t *testing.T) {
		const name, age = "Ray", 24
		assert.Equal(t, "Ray's age is 24", fmt.Sprintf("%s's age is %d", name, age))
	})

	// 在编译期间自定义函数均属于未知，因此无法用于常量的赋值，但内置函数可以使用
	t.Run("when declaring constants, we can use built-in functions", func(t *testing.T) {
		const strLength = len("string")
		assert.Equal(t, 6, strLength)
	})
}

// Go 并不支持 Java 中的 `enum` 关键字，但是我们可以用 `const` 关键字来达到相同的效果
func TestConstantsCanBeFactored(t *testing.T) {
	const (
		Sunday    = 0
		Monday    = 1
		Tuesday   = 2
		Wednesday = 3
		Thursday  = 4
		Friday    = 5
		Saturday  = 6
	)

	assert.Equal(t, "0 1 2 3 4 5 6", fmt.Sprint(Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday))
}

// 常量声明的时候可以声明类型，但这大多数时候都不是必须的，因为编译器可以自动推断
func TestConstantsDeclarationsWithType(t *testing.T) {
	const Pi float64 = 3.14
	const size int64 = 1024
	const u, v float32 = 0, 3
	assert.Equal(t, 3.14, Pi)
	assert.Equal(t, int64(1024), size)
	assert.Equal(t, float32(0), u)
	assert.Equal(t, float32(3), v)
}

// Go 有三个预定义的常量：`true`、`false`、`iota`
func TestPredefinedConstants(t *testing.T) {
	assert.True(t, true)
	assert.False(t, false)
	const a = iota
	assert.Equal(t, 0, a)
}

// 常量在缺省时会拥有与上一行的常量相同的值和类型
func TestEmptyConstantInGroups(t *testing.T) {
	const (
		a = 16
		b // same as a
		c = "char"
		d // same as d
	)
	assert.Equal(t, a, b)
	assert.Equal(t, c, d)
}

func TestIotaUsage(t *testing.T) {
	// `iota` 在每遇到一次常量声明时都会加 1，初始为 0
	t.Run("basic usage", func(t *testing.T) {
		const (
			a = iota // a == 0 (iota == 0)
			b = iota // b == 1 (iota == 1)
			c = iota // c == 2 (iota == 2)
		)
		assert.Equal(t, 0, a)
		assert.Equal(t, 1, b)
		assert.Equal(t, 2, c)
	})

	// 在常量按组声明时，`iota` 可以省略
	t.Run("omit iotas after first occurrence of iota", func(t *testing.T) {
		const (
			d = iota // d == 0 (iota == 0)
			e        // e == 1 (iota == 1)
			f        // f == 2 (iota == 2)
		)
		// 下面两行单独声明则都会为 0
		const v1 = iota
		const v2 = iota
		assert.Equal(t, 0, d)
		assert.Equal(t, 1, e)
		assert.Equal(t, 2, f)
		assert.Equal(t, 0, v1)
		assert.Equal(t, 0, v2)
	})

	// iota 只要有新的常量声明，就会保持自增，即使中间有新的常量声明也不会重置 iota 的值
	t.Run("iota will not reset when encounter explicit assignments", func(t *testing.T) {
		const (
			g = iota     // g == 0              (iota == 0)
			h            // h == 1              (iota == 1)
			i = "string" // i == "string"       (iota == 2)
			j            // j == i == "string"  (iota == 3)
			k = iota     // k == 4              (iota == 4)
		)

		assert.Equal(t, fmt.Sprintf("%s\n", "0 1 string string 4"), fmt.Sprintln(g, h, i, j, k))
	})

	// 只要常量组内使用了 `iota`，那么 `iota` 就会从第一行开始计数
	t.Run("iota will start counting from the first line", func(t *testing.T) {
		const (
			l = 7    // l == 7      (iota == 0)
			m = 8    // m == 8      (iota == 1)
			n        // n == m == 8 (iota == 2)
			o = iota // o == 3      (iota == 3)
			p        // p == 4      (iota == 4)
		)

		assert.Equal(t, "7 8 8 3 4", fmt.Sprint(l, m, n, o, p))
	})

	// 可以使用空白标识符 _ 来跳过一些不想使用的 iota 值
	t.Run("use blank identifier to discard values", func(t *testing.T) {
		const (
			q = iota // q = iota = 0
			r        // r = iota = 1
			_        // iota = 2
			_        // iota = 3
			s        // s = iota = 4
		)
		assert.Equal(t, "0 1 4", fmt.Sprint(q, r, s))
	})

	// 同一行定义多个包含 iota 的常量时，每个常量对应的 iota 值是相等的
	t.Run("multiple uses of iota in the same ConstSpec all have the same value", func(t *testing.T) {
		const (
			w       = iota             // w == 0           (iota == 0)
			x, y, z = iota, iota, iota // x == y == z == 1 (iota == 1)
		)

		assert.Equal(t, "0 1 1 1", fmt.Sprint(w, x, y, z))
	})
}

func TestExamples(t *testing.T) {
	t.Run("colors", func(t *testing.T) {
		// `type` 可以对一个已存在数据类型进行别名，这在某些时候更具可读性
		type Color int

		const (
			RED    Color = iota // RED == 0    (iota == 0)
			ORANGE              // ORANGE == 1 (iota == 1)
			YELLOW              // YELLOW == 2 (iota == 2)
			_                   //             (iota == 3)
			_                   //             (iota == 4)
			INDIGO              // INDIGO == 5 (iota == 5)
			VIOLET              // VIOLET == 6 (iota == 6)
		)

		assert.Equal(t, "0 1 2 5 6", fmt.Sprint(RED, ORANGE, YELLOW, INDIGO, VIOLET))

		var c Color
		assert.Equal(t, "constants.Color", fmt.Sprint(reflect.TypeOf(c)))
	})

	t.Run("size", func(t *testing.T) {
		type ByteSize float64

		const (
			_           = iota             // ignore 0
			KB ByteSize = 1 << (10 * iota) // 1 << (10*1)
			MB                             // 1 << (10*2)
			GB                             // 1 << (10*3)
			TB                             // 1 << (10*4)
			PB                             // 1 << (10*5)
			EB                             // 1 << (10*6)
			ZB                             // 1 << (10*7)
			YB                             // 1 << (10*8)
		)

		// 1024 1.048576e+06 1.073741824e+09 1.099511627776e+12 1.125899906842624e+15 1.152921504606847e+18 1.1805916207174113e+21 1.2089258196146292e+24
		assert.NotNil(t, fmt.Sprint(KB, MB, GB, TB, PB, EB, ZB, YB))
	})

	t.Run("bits and masks", func(t *testing.T) {
		const (
			bit0, mask0 = 1 << iota, 1<<iota - 1 // bit0 == 1, mask0 == 0  (iota == 0)
			bit1, mask1                          // bit1 == 2, mask1 == 1  (iota == 1)
			_, _                                 //                        (iota == 2, unused)
			bit3, mask3                          // bit3 == 8, mask3 == 7  (iota == 3)
		)
		assert.Equal(t, "1 0 2 1 8 7", fmt.Sprint(bit0, mask0, bit1, mask1, bit3, mask3))
	})
}
