package functions

// A closure is a function value that references variables from outside its body.
// The function may access and assign to the referenced variables;
// in this sense the function is "bound" to the variables.
import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCallbacksAcceptFuncAsAParam(t *testing.T) {
	add := func(a, b int) int {
		return a + b
	}

	callback := func(x, y int, f func(int, int) int) int {
		return f(x, y)
	}

	assert.Equal(t, 3, callback(1, 2, add))
}

func TestClosureCanAccessOutsideVariables(t *testing.T) {
	t.Run("example 1", func(t *testing.T) {
		multiplier := 2
		multiplyClosure := func(i int) int {
			return i * multiplier
		}

		total := 0
		for i := 0; i < 5; i++ {
			total += multiplyClosure(i)
		}

		assert.Equal(t, total, 20)
	})

	t.Run("example 2", func(t *testing.T) {
		j := 5
		f := func() func() (int, int) {
			i := 10
			return func() (int, int) {
				return i, j
			}
		}()

		r1, r2 := f()
		assert.Equal(t, r1, 10)
		assert.Equal(t, r2, 5)

		j *= 2
		r1, r2 = f()
		assert.Equal(t, r1, 10)
		assert.Equal(t, r2, 10)
	})
}

func TestAdder(t *testing.T) {
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
}

func TestSequentialAdder(t *testing.T) {
	sequentialAdder := func() func(int) int {
		var x int
		return func(i int) int {
			x += i
			return x
		}
	}

	var f = sequentialAdder()
	assert.Equal(t, 1, f(1))
	assert.Equal(t, 21, f(20))
	assert.Equal(t, 321, f(300))
}

func TestMakeAddSuffix(t *testing.T) {
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
}
