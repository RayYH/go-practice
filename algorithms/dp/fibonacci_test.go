package dp

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFibonacciNumbers(t *testing.T) {
	fibonacciNumbers := []int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377}

	for i := range fibonacciNumbers {
		assert.Equal(t, fibonacciNumbers[i], memoization(i))
		assert.Equal(t, fibonacciNumbers[i], tabulation(i))
	}
}
