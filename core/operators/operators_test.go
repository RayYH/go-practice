package main

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestArithmeticOperators(t *testing.T) {
	// x = q*y + r  and  |r| < |y|
	// 取余 % 符号始终与被除数一致
	t.Run("integer operators", func(t *testing.T) {
		assert.Equal(t, 8, 5+3)
		assert.Equal(t, 2, 5-3)
		assert.Equal(t, 1, 5/3)
		assert.Equal(t, 2, 5%3)

		assert.Equal(t, -2, -5+3)
		assert.Equal(t, -8, -5-3)
		assert.Equal(t, -1, -5/3)
		assert.Equal(t, -2, -5%3)

		assert.Equal(t, 2, 5+(-3))
		assert.Equal(t, 8, 5-(-3))
		assert.Equal(t, -1, 5/(-3))
		assert.Equal(t, 2, 5%(-3))

		assert.Equal(t, -8, -5+(-3))
		assert.Equal(t, -2, -5-(-3))
		assert.Equal(t, 1, -5/(-3))
		assert.Equal(t, -2, -5%(-3))

		assert.Equal(t, 128, math.MinInt8/-1)
	})

}
