package control_flows

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIf(t *testing.T) {
	t.Run("if", func(t *testing.T) {
		isGreaterThan := func(x, y int) bool {
			if x > y {
				return true
			}

			return false
		}
		assert.True(t, isGreaterThan(5, 4))
		assert.False(t, isGreaterThan(4, 4))
	})

	t.Run("if elseif else", func(t *testing.T) {
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
	})

	t.Run("if with a short assignment", func(t *testing.T) {
		isLessThanTen := func(x int) bool {
			// 使用简短方式 `:=` 声明的变量的作用域只存在于包括 else 块在内的 if 结构中
			if val := 10; val > x {
				return true
			} else {
				// 在 else 块中，val 是可见的
				assert.True(t, val < x)
				return false
			}
		}
		assert.True(t, isLessThanTen(9))
		assert.False(t, isLessThanTen(11))
	})
}

func TestSwitchCase(t *testing.T) {
	t.Run("using literals", func(t *testing.T) {
		// switch 结构，可以接受逗号分隔的一组字面量
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
	})

	t.Run("using expressions", func(t *testing.T) {
		compareToZero := func(x int) string {
			// 除了字面量外，Go 中 switch 的 case 还可以是一个表达式
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
	})

	t.Run("fallthrough instead of break", func(t *testing.T) {
		// Go 并不支持其他语言中的 `break` 关键字，在 Go 中，一旦成功地匹配到某个分支
		// 在执行完相应代码后就会退出整个 switch 代码块 (简直人性化)
		// 如果你需要继续执行下一个分支的语句，你必须显式使用 fallthrough 来达到继续往后匹配的目的
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
	})
}

// Go 中的循环只支持 `for` 关键字
func TestLoop(t *testing.T) {
	t.Run("c style loop", func(t *testing.T) {
		j := 0
		for i := 0; i < 5; i++ {
			assert.Equal(t, i, j)
			j++
		}
	})

	t.Run("for used as while", func(t *testing.T) {
		i, j := 0, 0
		for i < 5 {
			assert.Equal(t, i, j)
			i, j = i+1, j+1
		}
	})

	t.Run("infinite loop", func(t *testing.T) {
		for {
			got := rand.Intn(32)
			if got < 16 {
				assert.Equal(t, true, got < 32)
			} else {
				break
			}
		}
	})

	t.Run("for range", func(t *testing.T) {
		numbers := [4]int{1, 2, 3, 4}
		for index, value := range numbers {
			assert.Equal(t, index, value-1)
		}
	})
}

func TestJump(t *testing.T) {
	t.Run("break", func(t *testing.T) {
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
	})

	t.Run("continue", func(t *testing.T) {
		count := 0
		for i := 0; i < 10; i++ {
			if i == 5 {
				continue
			}
			assert.Equal(t, true, i != 5)
			count++
		}
		assert.Equal(t, 9, count)
	})

	t.Run("continue with label", func(t *testing.T) {
	LABEL:
		for i := 0; i <= 3; i++ {
			for j := 0; j <= 3; j++ {
				if j == 2 {
					continue LABEL
				}
				assert.True(t, j != 2)
			}
		}
	})

	t.Run("goto", func(t *testing.T) {
		i := 0
	HERE:
		i++
		if i == 5 {
			return
		} else {
			assert.Less(t, i, 5)
		}
		goto HERE
	})
}
