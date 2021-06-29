package operators

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestArithmeticOperators(t *testing.T) {
	t.Run("+ operator", func(t *testing.T) {
		assert.Equal(t, 8, 5+3)
	})

	t.Run("- operator", func(t *testing.T) {
		assert.Equal(t, 2, 5-3)
	})

	t.Run("* operator", func(t *testing.T) {
		assert.Equal(t, 15, 5*3)
	})

	t.Run("/ operator", func(t *testing.T) {
		assert.Equal(t, 1, 5/3)
		assert.Equal(t, -1, -5/3)
		assert.Equal(t, -1, 5/(-3))
		assert.Equal(t, 1, -5/(-3))
		// 浮点数除法，只要除数和被除数中有一个为浮点数，那么结果仍为浮点数
		a, b := 7.0, 3.0
		assert.Equal(t, a/b, 2.3333333333333335)
		a, b = 7.0, 3
		assert.Equal(t, a/b, 2.3333333333333335)
		// 只有两个类型相同的值才可以和二元运算符结合，Go 是强类型语言，因此不会进行隐式转换
		// 如果你使用 float64(7) / int(3.0)，编译器会报错
		// 但是如果你使用字面量 (常量) 来直接进行除法，编译是可以通过的
		assert.Equal(t, 7/3.0, 2.3333333333333335)
	})

	t.Run("% operator", func(t *testing.T) {
		// 取模运算符 % 只能作用于整数，余数的符号与被除数保持一致
		assert.Equal(t, 2, 5%3)
		assert.Equal(t, -2, (-5)%3)
		assert.Equal(t, 2, 5%(-3))
		assert.Equal(t, -2, (-5)%(-3))
	})
}

func TestRelationalOperators(t *testing.T) {
	// [0, 100) - 101 => [-101, -1)
	p := rand.Intn(100) - 101
	// [0, 100) + 101 => [101, 201)
	q := rand.Intn(100) + 101
	assert.False(t, p == q)
	assert.True(t, p != q)
	assert.True(t, p < q)
	assert.True(t, p <= q)
	assert.False(t, p > q)
	assert.False(t, p >= q)
}

func TestLogicalOperator(t *testing.T) {
	var p = rand.Intn(100) - 101
	var q = rand.Intn(100) + 101
	assert.True(t, p != q && p <= q)
	assert.True(t, p != q || p <= q)
	assert.True(t, !(p == q))
}

func TestBitwiseOperator(t *testing.T) {
	t.Run("bitwise AND", func(t *testing.T) {
		assert.Equal(t, 1, 1&1)
		assert.Equal(t, 0, 1&0)
		assert.Equal(t, 0, 0&1)
		assert.Equal(t, 0, 0&0)
		/*
		 * And
		 * 5 & 9
		 * 5 = 0000 0101
		 * 9 = 0000 1001
		 * -------------
		 *   = 0000 0001
		 *   = 1
		 */
		assert.Equal(t, 1, 5&9)
	})

	t.Run("bitwise OR", func(t *testing.T) {
		assert.Equal(t, 1, 1|1)
		assert.Equal(t, 1, 1|0)
		assert.Equal(t, 1, 0|1)
		assert.Equal(t, 0, 0|0)
		/*
		 * Or
		 * 5 | 9
		 * 5 = 0000 0101
		 * 9 = 0000 1001
		 * -------------
		 *   = 0000 1101
		 *   = 1 + 4 + 8
		 */
		assert.Equal(t, 1+4+8, 5|9)
	})

	t.Run("bitwise XOR", func(t *testing.T) {
		assert.Equal(t, 0, 1^1)
		assert.Equal(t, 1, 1^0)
		assert.Equal(t, 1, 0^1)
		assert.Equal(t, 0, 0^0)
		/*
		 * Xor
		 * 5 ^ 9
		 * 5 = 0000 0101
		 * 9 = 0000 1001
		 * -------------
		 *   = 0000 1100
		 *   = 4 + 8
		 */
		assert.Equal(t, 4+8, 5^9)
	})

	// https://stackoverflow.com/questions/34459450/what-is-the-operator-in-golang
	t.Run("bit clear (AND NOT)", func(t *testing.T) {
		// The C equivalent of the Go expression x &^ y is just x & ~y. That is literally "x AND (bitwise NOT of y)".
		// if you think of x | y as a way to turn on certain bits of x based on a mask constant y, then x &^ y is
		// doing the opposite and turns those same bits off (KEEP the different bits and clear the same bits)
		assert.Equal(t, 0, 1&^1) // 1 & 0 = 0
		assert.Equal(t, 1, 1&^0) // 1 & 1 = 1
		assert.Equal(t, 0, 0&^1) // 0 & 0 = 0
		assert.Equal(t, 0, 0&^0) // 0 & 1 = 0
		assert.Equal(t, 16, 0b00010100&^0b00001111)
		// 00010100
		// 00001111
		// --------
		// 00010000 = 16
	})

	t.Run("unary operator", func(t *testing.T) {
		/*
		 * In other programming languages, the bitwise complement operator is ~
		 * we should use ^ operator in Go.
		 *
		 * Complement
		 * 4
		 * 0000 0100 - original code
		 * 1111 1011 - inverse code
		 * 1111 1010 - minus one
		 * 1000 0101 - inverse all bits except sign bit
		 *
		 * There is a simple rule: A + (^A) = -1
		 */
		assert.Equal(t, -5, ^4)
		assert.Equal(t, -3, ^2)
	})

	t.Run("bitwise shift operator", func(t *testing.T) {
		assert.Equal(t, 2, 4>>1)
		assert.Equal(t, 1, 4>>2)
		// Keep moving will result 0
		assert.Equal(t, 0, 4>>3)
		assert.Equal(t, 0, 4>>4)

		assert.Equal(t, 8, 4<<1)
		assert.Equal(t, 16, 4<<2)
		assert.Equal(t, -2, -4>>1)
		assert.Equal(t, -1, -4>>2)
		// The sign will not be changed
		assert.Equal(t, -1, -4>>3)
		assert.Equal(t, -1, -4>>4)
		// -1 的补码是 11111111，所以右移不会有任何改变
		// 算术右移是带入符号位 (补 1)，逻辑右移是不带入符号位 (补 0)
		assert.Equal(t, -8, -4<<1)
		assert.Equal(t, -16, -4<<2)
	})
}

func TestAssignmentOperators(t *testing.T) {
	a := 1
	assert.Equal(t, 1, a)
	a += 1
	assert.Equal(t, 2, a)
	a -= 1
	assert.Equal(t, 1, a)
	a *= 3
	assert.Equal(t, 3, a)
	a /= 3
	assert.Equal(t, 1, a)

	b := 3
	b %= 2
	assert.Equal(t, 1, b)

	c := 0b00010100
	c &= 0b00000100
	assert.Equal(t, 4, c)
	c |= 0b00010000
	assert.Equal(t, 20, c)
	c ^= 0b00010101
	assert.Equal(t, 1, c)
}
