package main

import (
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

func checkValueAdvanced(x int) string {
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

func main() {

}
