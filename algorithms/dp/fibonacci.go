package dp

const maxSize = 100
const nil = -1

var memo [maxSize]int

func memoization(n int) int {
	if n > maxSize-1 {
		return nil
	}

	if memo[n] == nil {
		if n <= 1 {
			memo[n] = n
		} else {
			memo[n] = memoization(n-1) + memoization(n-2)
		}
	}

	return memo[n]
}

func tabulation(n int) int {
	tables := make([]int, n+1)

	if n <= 1 {
		return n
	}

	tables[0] = 0
	tables[1] = 1

	for i := 2; i <= n; i++ {
		tables[i] = tables[i-1] + tables[i-2]
	}

	return tables[n]
}

func init() {
	for i := range memo {
		memo[i] = nil
	}
}
