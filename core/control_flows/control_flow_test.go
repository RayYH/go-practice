package control_flows

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestIfStatement(t *testing.T) {
	isGreaterThan := func(x, y int) bool {
		if x > y {
			return true
		}

		return false
	}
	assert.True(t, isGreaterThan(5, 4))
	assert.False(t, isGreaterThan(4, 4))
}

func TestIfElseStatement(t *testing.T) {
	compareToZero := func(n int) string {
		if n < 0 {
			return "less than"
		} else if n == 0 {
			return "equal"
		} else {
			return "greater than"
		}
	}

	assert.Equal(t, "less than", compareToZero(-1))
	assert.Equal(t, "equal", compareToZero(0))
	assert.Equal(t, "greater than", compareToZero(1))
}

func TestIfWithAShortStatement(t *testing.T) {
	// Variables declared by the statement are only in scope until the end of the if.
	isLessThanTen := func(x int) bool {
		if val := 10; val > x {
			return true
		} else {
			// Variables declared inside an if short statement are also available inside any of the else blocks.
			assert.NotNil(t, val)
			return false
		}
	}
	assert.True(t, isLessThanTen(9))
	assert.False(t, isLessThanTen(11))
}

func TestSwitchCaseUsingLiterals(t *testing.T) {
	checkValue := func(x int) string {
		switch x {
		case 98, 99:
			return "98 or 99"
		case 100:
			return "100"
		default:
			return "< 98 or > 100"
		}
	}
	assert.Equal(t, "98 or 99", checkValue(98))
	assert.Equal(t, "98 or 99", checkValue(99))
	assert.Equal(t, "100", checkValue(100))
	assert.Equal(t, "< 98 or > 100", checkValue(97))
	assert.Equal(t, "< 98 or > 100", checkValue(101))
}

func TestSwitchCaseUsingExpressions(t *testing.T) {
	// This construct can be a clean way to write long if-then-else chains.
	compareToZero := func(x int) string {
		switch {
		case x < 0:
			return "<"
		case x == 0:
			return "=="
		default:
			return ">"
		}
	}
	assert.Equal(t, "<", compareToZero(-1))
	assert.Equal(t, "==", compareToZero(0))
	assert.Equal(t, ">", compareToZero(1))
}

func TestFallThroughInsideSwitchCase(t *testing.T) {
	minusOneOrTwo := func(x int) int {
		switch x {
		case 0:
			fallthrough
		case 1:
			fallthrough
		case 2:
			return 2
		}

		return -1
	}
	assert.Equal(t, 2, minusOneOrTwo(0))
	assert.Equal(t, 2, minusOneOrTwo(1))
	assert.Equal(t, 2, minusOneOrTwo(2))
	assert.Equal(t, -1, minusOneOrTwo(3))
}

func TestCStyleForLoop(t *testing.T) {
	j := 0
	for i := 0; i < 5; i++ {
		assert.Equal(t, i, j)
		j++
	}
}

func TestForUsedAsWhileLoop(t *testing.T) {
	i := 0
	j := 0
	for i < 5 {
		assert.Equal(t, i, j)
		i += 1
		j += 1
	}
}

func TestInfiniteLoop(t *testing.T) {
	for {
		got := rand.Intn(32)
		if got < 16 {
			assert.Equal(t, true, got < 32)
		} else {
			break
		}
	}
}

func TestForRangeLoop(t *testing.T) {
	numbers := [4]int{1, 2, 3, 4}
	for index, value := range numbers {
		assert.Equal(t, index, value-1)
	}
}

func TestBreakStatement(t *testing.T) {
	count := 0
	for i := 0; i < 3; i++ {
		if i == 2 {
			break
		}
		for j := 0; j < 10; j++ {
			if j == 5 {
				break
			}
			// i = 0, j =0 1 2 3 4
			// i = 1, j =0 1 2 3 4
			count++
			assert.Equal(t, true, j != 5)
		}
	}
	assert.Equal(t, 10, count)
}

func TestContinueStatement(t *testing.T) {
	count := 0
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		assert.Equal(t, true, i != 5)
		count++
	}
	assert.Equal(t, 9, count)
}

func TestLabelForContinueOrBreakStatements(t *testing.T) {
LABEL:
	for i := 0; i <= 3; i++ {
		for j := 0; j <= 3; j++ {
			if j == 2 {
				continue LABEL
			}
			assert.Equal(t, true, j != 2)
		}
	}
}

func TestGotoStatementWithLabel(t *testing.T) {
	i := 0
HERE:
	i++
	if i == 5 {
		return
	} else {
		assert.Less(t, i, 5)
	}
	goto HERE
}
