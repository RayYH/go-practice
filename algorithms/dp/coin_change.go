package dp

import (
	"math"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minimumCoinsMemoization(coins []int, amount int) int {
	memo := make(map[int]int, amount+1)
	return minimumCoinsMemo(coins, amount, memo)
}

func minimumCoinsMemo(coins []int, amount int, memo map[int]int) int {
	// check memo
	ans, ok := memo[amount]
	if ok {
		return ans
	}

	// base case and param validation
	if amount < 0 {
		return -1
	}

	if amount == 0 {
		return 0
	}

	// update memo via recursive strategy (check if no solution)
	res := math.MaxInt32
	for i := 0; i <= amount; i++ {
		for _, v := range coins {
			subSolution := minimumCoinsMemo(coins, amount-v, memo)
			if subSolution == -1 {
				continue
			}
			res = min(res, subSolution+1)
		}
	}

	// return solution stored in memo (when initial value not changed means no solution)
	if res == math.MaxInt32 {
		memo[amount] = -1
	} else {
		memo[amount] = res
	}
	return memo[amount]
}

func minimumCoinsTabulation(coins []int, amount int) int {
	// param validation
	if amount < 0 {
		return -1
	}

	// init tables
	tables := make([]int, amount+1)
	for i := range tables {
		tables[i] = math.MaxInt32
	}

	// base case in tables
	tables[0] = 0

	// update tables via recursive strategy (no solution check)
	for i := 0; i <= amount; i++ {
		for _, coin := range coins {
			if i-coin < 0 {
				continue
			}
			subSolution := tables[i-coin]
			tables[i] = min(tables[i], subSolution+1)
		}
	}

	// return solution (if initial value has not been modified, then means no solution)
	if tables[amount] == math.MaxInt32 {
		return -1
	}

	return tables[amount]
}

func totalCombinationsTabulation(coins []int, amount int) int {
	// params validation
	if amount < 0 {
		return 0
	}

	// init tables
	tables := make([]int, amount+1)

	// base case
	tables[0] = 1

	// update tables
	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			tables[i] += tables[i-coin]
		}
	}

	// return solution stored in tables
	return tables[amount]
}
