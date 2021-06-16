package functions

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// A defer statement defers the execution of a function until the surrounding function returns.
// The deferred call's arguments are evaluated immediately, but the function call
// is not executed until the surrounding function returns.
func Example_greet() {
	greet := func() {
		fmt.Println("Hello")
	}
	defer greet()
	fmt.Println("DEFER")
	// Output:
	// DEFER
	// Hello
}

func TestGetValue(t *testing.T) {
	// Deferred function calls are pushed onto a stack.
	// When a function returns, its deferred calls are executed in last-in-first-out order.
	getValue := func() (ret int) {
		defer func() {
			ret++
		}()
		return 1
	}

	assert.Equal(t, getValue(), 2)
}

func Example_deferOrders() {
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
	// Output:
	// 3
	// 2
	// 1
}
