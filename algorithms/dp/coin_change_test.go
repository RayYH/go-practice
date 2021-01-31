package dp

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type testCase struct {
	coins  []int
	amount int
	ans    int
}

func TestMinimumCoins(t *testing.T) {
	testCases := make([]testCase, 5)
	testCases[0] = testCase{
		coins:  []int{1, 2, 5},
		amount: 11,
		ans:    3,
	}
	testCases[1] = testCase{
		coins:  []int{1, 2, 5},
		amount: 10,
		ans:    2,
	}
	testCases[2] = testCase{
		coins:  []int{1, 2, 5},
		amount: 30,
		ans:    6,
	}
	testCases[3] = testCase{
		coins:  []int{1, 2, 5},
		amount: -1,
		ans:    -1,
	}
	testCases[4] = testCase{
		coins:  []int{7, 8, 9},
		amount: 4,
		ans:    -1,
	}

	for _, testCase := range testCases {
		assert.Equal(t, minimumCoinsMemoization(testCase.coins, testCase.amount), testCase.ans)
		assert.Equal(t, minimumCoinsTabulation(testCase.coins, testCase.amount), testCase.ans)
	}
}

func TestTotalCombinationsTabulation(t *testing.T) {
	testCases := make([]testCase, 3)
	testCases[0] = testCase{
		coins:  []int{1, 2, 5},
		amount: 5,
		ans:    4,
	}
	testCases[1] = testCase{
		coins:  []int{2},
		amount: 3,
		ans:    0,
	}
	testCases[2] = testCase{
		coins:  []int{10},
		amount: 10,
		ans:    1,
	}

	for _, testCase := range testCases {
		assert.Equal(t, totalCombinationsTabulation(testCase.coins, testCase.amount), testCase.ans)
	}
}
