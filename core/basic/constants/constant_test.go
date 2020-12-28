package constants

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPi(t *testing.T) {
	assert.Equal(t, 3.14159, Pi, "Pi should be equal to 3.14159")
}

func TestDays(t *testing.T) {
	assert.Equal(t, "0 1 2 3 4 5 6", fmt.Sprint(Sunday, Monday, Tuesday, Wednesday, Thursday, Friday, Saturday))
}

func TestPersonInfo(t *testing.T) {
	assert.Equal(t, "Ray's age is 24", fmt.Sprintf("%s's age is %d", name, age))
}

func TestIotaUsageOne(t *testing.T) {
	assert.Equal(t, 0, a)
	assert.Equal(t, 1, b)
	assert.Equal(t, 2, c)
}

func TestIotaUsageTwo(t *testing.T) {
	assert.Equal(t, 0, d)
	assert.Equal(t, 1, e)
	assert.Equal(t, 2, f)
}

func TestIotaUsageThree(t *testing.T) {
	assert.Equal(t, fmt.Sprintf("%s\n", "0 1 string string 4"), fmt.Sprintln(g, h, i, j, k))
}

func TestIotaUsageFour(t *testing.T) {
	assert.Equal(t, "7 8 8 3 4", fmt.Sprint(l, m, n, o, p))
}

func TestIotaUsageFive(t *testing.T) {
	assert.Equal(t, "0 1 2 5 6", fmt.Sprint(RED, ORANGE, YELLOW, INDIGO, VIOLET))
}

func TestIotaUsageSix(t *testing.T) {
	assert.Equal(t, "0 1 1 1", fmt.Sprint(w, x, y, z))
}

func ExampleDisplaySizes() {
	DisplaySizes()
	// Output: sizes: 1024 1.048576e+06 1.073741824e+09 1.099511627776e+12 1.125899906842624e+15 1.152921504606847e+18 1.1805916207174113e+21 1.2089258196146292e+24
}

func TestDeclareConstInsideAFunc(t *testing.T) {
	const Truth = true
	assert.True(t, Truth)
}

func TestConstantWillNotLosePrecision(t *testing.T) {
	assert.Equal(t, float32(3.1415927), float32(HigherPrecisionPi))
	// default is float64
	assert.Equal(t, 3.141592653589793, HigherPrecisionPi)
	assert.NotEqual(t, LessThanOne, 3.141592653589793/3.141592653589793)
}

func TestBitsAndMasks(t *testing.T) {
	assert.Equal(t, "1 0 2 1 8 7", fmt.Sprint(bit0, mask0, bit1, mask1, bit3, mask3))
}
