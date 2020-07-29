package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func getTrue() bool {
	return true
}

func TestIfStatement(t *testing.T) {
	assert.True(t, getTrue(), "getTrue should returns True")
}

func TestIsGreater(t *testing.T) {
	assert.True(t, isGreater(10, 5), "10 should be greater than 5")
}

func TestLessThanTen(t *testing.T) {
	assert.False(t, lessThanTen(11), "11 should be greater than 10")
}

func TestMySqrt(t *testing.T) {
	_, ok := mySqrt(25.0)
	if !ok {
		t.Errorf("sqrt(25.0) is invalid")
	}

	_, ok = mySqrt(-1)

	if ok {
		t.Errorf("sqrt(-1) is valid")
	}
}

func TestCheckValue(t *testing.T) {
	assert.Equal(t, "100", checkValue(100), "100 == 100")
}

func TestCheckValueUseIf(t *testing.T) {
	assert.Equal(t, "100", checkValueUseIf(100), "100 == 100")
}

func TestCheckValueExpression(t *testing.T) {
	assert.Equal(t, "x > 0", checkValueUseExpression(1), "1 should be greater than 0")
}

func TestFallThroughExample(t *testing.T) {
	assert.Equal(t, 2, fallThroughExample(0), "fallthrough will continue execution")
}

func ExampleShowForStatement() {
	for i := 0; i < 5; i++ {
		fmt.Printf("This is the %d iteration\n", i)
	}
	// Output:
	// This is the 0 iteration
	// This is the 1 iteration
	// This is the 2 iteration
	// This is the 3 iteration
	// This is the 4 iteration
}

func ExampleUseForAsWhile() {
	i := 5
	for i > 0 {
		fmt.Printf("i = %d\n", i)
		i--
	}

	// Output:
	// i = 5
	// i = 4
	// i = 3
	// i = 2
	// i = 1
}

func ExampleForInfiniteLoop() {
	for {
		fmt.Print("start...")
		break
	}
	// Output:
	// start...
}

func ExampleForRange() {
	nums := []int{1, 2, 3, 4}
	var num int
	for _, num = range nums {
		fmt.Printf("%d ", num)
	}
	// Output:
	// 1 2 3 4
}

func ExampleBreak() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 10; j++ {
			if j > 5 {
				break
			}
			fmt.Print(j)
		}
		fmt.Print("  ")
	}
	// Output:
	// 012345  012345  012345
}

func ExampleContinue() {
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Print(i)
		fmt.Print(" ")
	}
	// Output:
	// 0 1 2 3 4 6 7 8 9
}

func ExampleUseLabel() {
LABEL:
	for i := 0; i <= 3; i++ {
		for j := 0; j <= 3; j++ {
			if j == 2 {
				continue LABEL
			}
			fmt.Printf("i is: %d, and j is: %d\n", i, j)
		}
	}

	// Output:
	// i is: 0, and j is: 0
	// i is: 0, and j is: 1
	// i is: 1, and j is: 0
	// i is: 1, and j is: 1
	// i is: 2, and j is: 0
	// i is: 2, and j is: 1
	// i is: 3, and j is: 0
	// i is: 3, and j is: 1
}

func ExampleGotoStatement() {
	i := 0
HERE:
	fmt.Print(i)
	i++
	if i == 5 {
		return
	}
	goto HERE
	// Output:
	// 01234
}
