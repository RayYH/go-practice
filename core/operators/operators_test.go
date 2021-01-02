package operators

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
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
	})

	t.Run("% operator", func(t *testing.T) {
		// % doesn't work for floats
		// x = q*y + r  and  |r| < |y|
		// given c = a % b, the sign of c to be the sign of a
		assert.Equal(t, 2, 5%3)
		assert.Equal(t, -2, -5%3)
		assert.Equal(t, 2, 5%(-3))
		assert.Equal(t, -2, -5%(-3))
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

	// https://stackoverflow.com/questions/34459450/what-is-the-operator-in-golang
	t.Run("bit clear (AND NOT)", func(t *testing.T) {
		// The C equivalent of the Go expression x &^ y is just x & ~y. That is literally "x AND (bitwise NOT of y)".
		// if you think of x | y as a way to turn on certain bits of x based on a mask constant y, then x &^ y is
		// doing the opposite and turns those same bits off
		// KEEP the different bits and clear the same bits
		assert.Equal(t, 1&^1, 0) // 1 & 0 = 0
		assert.Equal(t, 1&^0, 1) // 1 & 1 = 1
		assert.Equal(t, 0&^1, 0) // 0 & 0 = 0
		assert.Equal(t, 0&^0, 0) // 0 & 1 = 0
		assert.Equal(t, 0b00010100&^0b00001111, 16)
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
		assert.Equal(t, ^4, -5)
		assert.Equal(t, ^2, -3)
		assert.Equal(t, 4>>1, 2)
		assert.Equal(t, 4>>2, 1)
		// Keep moving will result 0
		assert.Equal(t, 4>>3, 0)
		assert.Equal(t, 4>>4, 0)
		assert.Equal(t, 4<<1, 8)
		assert.Equal(t, 4<<2, 16)
		assert.Equal(t, -4>>1, -2)
		assert.Equal(t, -4>>2, -1)
		// The sign will not be changed
		assert.Equal(t, -4>>3, -1)
		assert.Equal(t, -4>>4, -1)
		assert.Equal(t, -4<<1, -8)
		assert.Equal(t, -4<<2, -16)
	})
}

// other assignment operators：=, +=, -=, *=, /=, %=, &=, ^=, |=
