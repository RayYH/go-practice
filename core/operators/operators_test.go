package main

import (
	"github.com/stretchr/testify/assert"
	"math"
	"math/rand"
	"testing"
)

func TestArithmeticOperators(t *testing.T) {
	// x = q*y + r  and  |r| < |y|
	// 取余 % 符号始终与被除数一致
	t.Run("integer operators", func(t *testing.T) {
		assert.Equal(t, 8, 5+3)
		assert.Equal(t, 2, 5-3)
		assert.Equal(t, 15, 5*3)
		assert.Equal(t, 1, 5/3)
		assert.Equal(t, 2, 5%3)

		assert.Equal(t, -2, -5+3)
		assert.Equal(t, -8, -5-3)
		assert.Equal(t, -15, -5*3)
		assert.Equal(t, -1, -5/3)
		assert.Equal(t, -2, -5%3)

		assert.Equal(t, 2, 5+(-3))
		assert.Equal(t, 8, 5-(-3))
		assert.Equal(t, -15, 5*(-3))
		assert.Equal(t, -1, 5/(-3))
		assert.Equal(t, 2, 5%(-3))

		assert.Equal(t, -8, -5+(-3))
		assert.Equal(t, -2, -5-(-3))
		assert.Equal(t, 15, -5*(-3))
		assert.Equal(t, 1, -5/(-3))
		assert.Equal(t, -2, -5%(-3))

		assert.Equal(t, 128, math.MinInt8/-1)
	})

}

func TestRelationalOperators(t *testing.T) {
	t.Run("relational operators", func(t *testing.T) {
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
	})
}

func TestLogicalOperator(t *testing.T) {
	t.Run("logical operator", func(t *testing.T) {
		var p = rand.Intn(100) - 101
		var q = rand.Intn(100) + 101
		assert.True(t, p != q && p <= q)
		assert.True(t, p != q || p <= q)
		assert.True(t, !(p == q))
	})
}

func TestBitwiseOperator(t *testing.T) {
	t.Run("bitwise AND", func(t *testing.T) {
		assert.Equal(t, 1&1, 1)
		assert.Equal(t, 1&0, 0)
		assert.Equal(t, 0&1, 0)
		assert.Equal(t, 0&0, 0)
		/*
		 * And
		 * 5 & 9
		 * 5 = 0000 0101
		 * 9 = 0000 1001
		 * -------------
		 *   = 0000 0001
		 *   = 1
		 */
		assert.Equal(t, 5&9, 1)
	})
	t.Run("bitwise OR", func(t *testing.T) {
		assert.Equal(t, 1|1, 1)
		assert.Equal(t, 1|0, 1)
		assert.Equal(t, 0|1, 1)
		assert.Equal(t, 0|0, 0)
		/*
		 * Or
		 *  5 | 9
		 * 5 = 0000 0101
		 * 9 = 0000 1001
		 * -------------
		 *   = 0000 1101
		 *   = 1 + 4 + 8
		 */
		assert.Equal(t, 5|9, 1+4+8)
	})
	t.Run("bitwise XOR", func(t *testing.T) {
		assert.Equal(t, 1^1, 0)
		assert.Equal(t, 1^0, 1)
		assert.Equal(t, 0^1, 1)
		assert.Equal(t, 0^0, 0)
		/*
		 * Xor
		 * 5 ^ 9
		 * 5 = 0000 0101
		 * 9 = 0000 1001
		 * -------------
		 *   = 0000 1100
		 *   = 4 + 8
		 */
		assert.Equal(t, 5^9, 4+8)
	})
	t.Run("bit clear (AND NOT)", func(t *testing.T) {
		// The C equivalent of the Go expression x &^ y is just x & ~y. That is literally "x AND (bitwise NOT of y)".
		// https://stackoverflow.com/questions/34459450/what-is-the-operator-in-golang
		// opposite to |
		// 将运算符左边数据相异的位保留，相同位清零
		assert.Equal(t, 1&^1, 0)
		assert.Equal(t, 1&^0, 1)
		assert.Equal(t, 0&^1, 0)
		assert.Equal(t, 0&^0, 0)
		assert.Equal(t, 0b00010100&^0b00001111, 16)
	})
	t.Run("unary operator", func(t *testing.T) {
		/*
		 * Complement
		 * 4
		 * 0000 0100 - 原码 (正数的原码即为真值)
		 * 1111 1011 - 取反，得到的结果在计算机中存储的形式是补码，此补码指示其值是一个负数 (最高位是1)
		 * 1111 1010 - -1 (负数的补码由其原码取反 +1 而得，因此我们需要 -1 再对符号位之外的各位取反)
		 * 1000 0101 - 取反 (1 + 4)
		 *
		 * There is a simple rule: A + (~A) = -1
		 */
		assert.Equal(t, ^4, -5)
		assert.Equal(t, ^2, -3)
		assert.Equal(t, 4>>1, 2)
		assert.Equal(t, 4>>2, 1)
		// 继续右移会变成 0
		assert.Equal(t, 4>>3, 0)
		assert.Equal(t, 4>>4, 0)
		assert.Equal(t, 4<<1, 8)
		assert.Equal(t, 4<<2, 16)
		assert.Equal(t, -4>>1, -2)
		assert.Equal(t, -4>>2, -1)
		// 符号不会改变
		assert.Equal(t, -4>>3, -1)
		assert.Equal(t, -4>>4, -1)
		assert.Equal(t, -4<<1, -8)
		assert.Equal(t, -4<<2, -16)
	})
}

// 这里不再对赋值运算符举例说明：=、+=、-=、*=、/=、%=、&=、^=、|=
