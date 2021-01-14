package functions

// A closure is a function value that references variables from outside its body.
// The function may access and assign to the referenced variables;
// in this sense the function is "bound" to the variables.
import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestCallback(t *testing.T) {
	add := func(a, b int) int {
		return a + b
	}

	// callbacks accept func as a param
	callback := func(x, y int, f func(int, int) int) int {
		return f(x, y)
	}
	assert.Equal(t, 3, callback(1, 2, add))
}

func TestClosureCanAccessOutsideVariables(t *testing.T) {
	multiplier := 2
	multiplyClosure := func(i int) int {
		return i * multiplier
	}

	total := 0
	for i := 0; i < 5; i++ {
		total += multiplyClosure(i)
	}

	assert.Equal(t, total, 20)
}

func addTwo() func(int) int {
	return func(b int) int {
		return b + 2
	}
}

func adder(a int) func(int) int {
	return func(b int) int {
		return a + b
	}
}

func TestAdder(t *testing.T) {
	plusTwo := addTwo()
	assert.Equal(t, 5, plusTwo(3))
	plus := adder(4)
	assert.Equal(t, 9, plus(5))
}

func sequentialAdder() func(int) int {
	var x int
	return func(i int) int {
		x += i
		return x
	}
}

func TestSequentialAdder(t *testing.T) {
	var f = sequentialAdder()
	assert.Equal(t, 1, f(1))
	assert.Equal(t, 21, f(20))
	assert.Equal(t, 321, f(300))
}

func makeAddSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func TestMakeAddSuffix(t *testing.T) {
	addBmp := makeAddSuffix(".bmp")
	addJpeg := makeAddSuffix(".jpeg")
	assert.Equal(t, "file.bmp", addBmp("file"))
	assert.Equal(t, "file.jpeg", addJpeg("file"))
}
