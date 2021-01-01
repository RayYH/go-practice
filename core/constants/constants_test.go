package constants

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDeclareConstInsideAFunc(t *testing.T) {
	const Truth = true
	assert.True(t, Truth)
}

func TestSimpleDeclaration(t *testing.T) {
	const Pi = 3.14159
	assert.Equal(t, 3.14159, Pi)
}

func TestFactoringConstants(t *testing.T) {
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

func TestMultipleConstantsAtOneLine(t *testing.T) {
	const name, age = "Ray", 24
	assert.Equal(t, "Ray's age is 24", fmt.Sprintf("%s's age is %d", name, age))
}

func TestSimpleIotaExample(t *testing.T) {
	const (
		a = iota // a == 0 (iota == 0)
		b = iota // b == 1 (iota == 1)
		c = iota // c == 2 (iota == 2)
	)
	assert.Equal(t, 0, a)
	assert.Equal(t, 1, b)
	assert.Equal(t, 2, c)
}

// we can omit iota after iota's first occurrence
func TestOmitIotasAfterFirstOccurrenceOfIota(t *testing.T) {
	const (
		d = iota // d == 0 (iota == 0)
		e        // e == 1 (iota == 1)
		f        // f == 2 (iota == 2)
	)
	assert.Equal(t, 0, d)
	assert.Equal(t, 1, e)
	assert.Equal(t, 2, f)
}

// explicitly assignments will not reset the iota value
func TestIotaWillNotBeResetWhenEncounterExplicitlyAssignments(t *testing.T) {
	const (
		g = iota     // g == 0              (iota == 0)
		h            // h == 1              (iota == 1)
		i = "string" // i == "string"       (iota == 2)
		j            // j == i == "string"  (iota == 3)
		k = iota     // k == 4              (iota == 4)
	)

	assert.Equal(t, fmt.Sprintf("%s\n", "0 1 string string 4"), fmt.Sprintln(g, h, i, j, k))
}

func TestIotaWillStartCountingFormTheFirstLine(t *testing.T) {
	const (
		l = 7    // l == 7      (iota == 0)
		m = 8    // m == 8      (iota == 1)
		n        // n == m == 8 (iota == 2)
		o = iota // o == 3      (iota == 3)
		p        // p == 4      (iota == 4)
	)

	assert.Equal(t, "7 8 8 3 4", fmt.Sprint(l, m, n, o, p))
}

// By definition, multiple uses of iota in the same ConstSpec all have the same value
func TestIotaAtTheSameLine(t *testing.T) {
	const (
		w       = iota             // w == 0           (iota == 0)
		x, y, z = iota, iota, iota // x == y == z == 1 (iota == 1)
	)

	assert.Equal(t, "0 1 1 1", fmt.Sprint(w, x, y, z))
}

func TestColorConstantsViaIota(t *testing.T) {
	// use type to alias a known type
	type Color int

	// The blank identifier can be assigned or declared with any value of any type, with the value discarded harmlessly.
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
}

func TestSizeConstantsViaIota(t *testing.T) {
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

	assert.Equal(t, "1024 1.048576e+06 1.073741824e+09 1.099511627776e+12 1.125899906842624e+15 1.152921504606847e+18 1.1805916207174113e+21 1.2089258196146292e+24", fmt.Sprint(KB, MB, GB, TB, PB, EB, ZB, YB))
}

func TestConstantWillNotLosePrecision(t *testing.T) {
	// Numeric constants represent exact values of arbitrary precision and do not overflow.
	const HigherPrecisionPi = 3.14159265358979323846264338327950288419716939937510582097494459
	const LessThanOne = 3.141592653589793 / HigherPrecisionPi
	assert.Equal(t, float32(3.1415927), float32(HigherPrecisionPi))
	// default is float64
	assert.Equal(t, 3.141592653589793, HigherPrecisionPi)
	assert.NotEqual(t, LessThanOne, 3.141592653589793/3.141592653589793)
}

func TestBitsAndMasksViaSkippingIota(t *testing.T) {
	const (
		bit0, mask0 = 1 << iota, 1<<iota - 1 // bit0 == 1, mask0 == 0  (iota == 0)
		bit1, mask1                          // bit1 == 2, mask1 == 1  (iota == 1)
		_, _                                 //                        (iota == 2, unused)
		bit3, mask3                          // bit3 == 8, mask3 == 7  (iota == 3)
	)
	assert.Equal(t, "1 0 2 1 8 7", fmt.Sprint(bit0, mask0, bit1, mask1, bit3, mask3))
}
