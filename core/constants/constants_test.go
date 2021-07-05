package constants

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeclarations(t *testing.T) {
	t.Run("define constants via const keyword", func(t *testing.T) {
		t.Run("with specified types", func(t *testing.T) {
			const Pi float64 = 3.14
			const size int64 = 1024
			assert.Equal(t, 3.14, Pi)
			assert.Equal(t, int64(1024), size)
		})

		t.Run("without specified types", func(t *testing.T) {
			const Pi = 3.14159
			assert.Equal(t, 3.14159, Pi)
		})
	})

	t.Run("constants can be declared at one line", func(t *testing.T) {
		t.Run("with specified types", func(t *testing.T) {
			const u, v float32 = 0, 3
			assert.Equal(t, float32(0), u)
			assert.Equal(t, float32(3), v)
		})

		t.Run("without specified types", func(t *testing.T) {
			const name, age = "Ray", 24
			assert.Equal(t, "Ray's age is 24", fmt.Sprintf("%s's age is %d", name, age))
		})
	})

	t.Run("when declaring constants, we can use built-in functions", func(t *testing.T) {
		const strLength = len("string")
		assert.Equal(t, 6, strLength)
	})
}

func TestConstantsCanBeFactoredIntoBlocks(t *testing.T) {
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

func TestPredefinedConstants(t *testing.T) {
	assert.True(t, true)
	assert.False(t, false)
	const a = iota
	assert.Equal(t, 0, a)
}

func TestConstantWithoutGivenValueInGroups(t *testing.T) {
	const (
		a = 16
		b
		c = "char"
		d
		e
	)
	assert.Equal(t, a, b)
	assert.Equal(t, c, d)
	assert.Equal(t, d, e)
}

func TestIotaUsage(t *testing.T) {
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

	t.Run("omit iotas after first occurrence of iota", func(t *testing.T) {
		const (
			d = iota // d == 0 (iota == 0)
			e        // e == 1 (iota == 1)
			f        // f == 2 (iota == 2)
		)
		assert.Equal(t, 0, d)
		assert.Equal(t, 1, e)
		assert.Equal(t, 2, f)

		const v1 = iota
		const v2 = iota
		assert.Equal(t, 0, v1)
		assert.Equal(t, 0, v2)
	})

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
