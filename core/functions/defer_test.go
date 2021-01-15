package functions

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func greet() {
	fmt.Println("Hello")
}

// A defer statement defers the execution of a function until the surrounding function returns.
// The deferred call's arguments are evaluated immediately, but the function call
// is not executed until the surrounding function returns.
func Example_greet() {
	defer greet()
	fmt.Println("DEFER")
	// Output:
	// DEFER
	// Hello
}

// Deferred function calls are pushed onto a stack.
// When a function returns, its deferred calls are executed in last-in-first-out order.
func getValue() (ret int) {
	defer func() {
		ret++
	}()
	return 1
}

func TestGetValue(t *testing.T) {
	assert.Equal(t, getValue(), 2)
}
