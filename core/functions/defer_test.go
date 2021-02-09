package functions

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func greet() {
	fmt.Println("Hello")
}

func print(msg string) {
	fmt.Print(msg)
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

// FILO
func Example_print() {
	defer print("1")
	defer print("2")
	defer print("3")
	// Output:
	// 321
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
