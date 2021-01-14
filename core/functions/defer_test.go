package functions

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func greet() {
	fmt.Println("Hello")
}

// function invocation can be deferred
func Example_greet() {
	defer greet()
	fmt.Println("DEFER")
	// Output:
	// DEFER
	// Hello
}

// A defer statement pushes a function call onto a list.
// The list of saved calls is executed after the surrounding function returns.
func getValue() (ret int) {
	defer func() {
		ret++
	}()
	return 1
}

func TestGetValue(t *testing.T) {
	assert.Equal(t, getValue(), 2)
}
