package control_flows

import (
	"fmt"
	"math"
)

func isGreater(x, y int) bool {
	if x > y {
		return true
	}

	return false
}

func lessThanTen(x int) bool {
	if val := 10; val > x {
		return true
	}

	return false
}

func mySqrt(f float64) (v float64, ok bool) {
	if f < 0 {
		return
	}

	return math.Sqrt(f), true
}

func checkValueUseIf(x int) string {
	if x == 98 || x == 99 {
		return "98 or 99"
	} else if x == 100 {
		return "100"
	} else {
		return "< 98 or > 100"
	}
}

func checkValue(x int) string {
	switch x {
	case 98, 99:
		return "98 or 99"
	case 100:
		return "100"
	default:
		return "< 98 or > 100"
	}
}

func checkValueUseExpression(x int) string {
	switch {
	case x == 0:
		return "x == 0"
	case x > 0:
		return "x > 0"
	default:
		return "x < 0"
	}
}

func fallThroughExample(x int) int {
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

func ShowForStatement() {
	for i := 0; i < 5; i++ {
		fmt.Printf("This is the %d iteration\n", i)
	}
}

func UseForAsWhile() {
	i := 5
	for i > 0 {
		fmt.Printf("i = %d\n", i)
		i--
	}
}

func ForInfiniteLoop() {
	for {
		fmt.Print("start...")
		break
	}
}

func ForRange() {
	nums := []int{1, 2, 3, 4}
	var num int
	for _, num = range nums {
		fmt.Printf("%d ", num)
	}
}

func Break() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 10; j++ {
			if j > 5 {
				break
			}
			fmt.Print(j)
		}
		fmt.Print("  ")
	}
}

func Continue() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Print(i)
		fmt.Print(" ")
	}
}

func UseLabel() {
LABEL:
	for i := 0; i <= 3; i++ {
		for j := 0; j <= 3; j++ {
			if j == 2 {
				continue LABEL
			}
			fmt.Printf("i is: %d, and j is: %d\n", i, j)
		}
	}
}

func GotoStatement() {
	i := 0
HERE:
	fmt.Print(i)
	i++
	if i == 5 {
		return
	}
	goto HERE
}
