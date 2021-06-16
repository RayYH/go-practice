package errors

import (
	"errors"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBasicGolangExceptionModel(t *testing.T) {
	adder := func(a, b int) (ret int, err error) {
		if a < 0 || b < 0 {
			err = errors.New("should be non-negative numbers")
			return
		}

		return a + b, nil
	}
	c, _ := adder(4, 5)
	assert.Equal(t, c, 9)
	d, e := adder(-1, 2)
	assert.Equal(t, d, 0)
	assert.Error(t, e)
}

func TestMyError(t *testing.T) {
	run := func() error {
		return &MyError{
			"now",
			"something happened",
		}
	}
	err := run()
	assert.NotNil(t, err)
	assert.Equal(t, "at now, something happened", fmt.Sprint(err))
}

// https://golangbot.com/panic-and-recover/
// recover is a builtin function that is used to regain control of a panicking program.
func recoverFullName() {
	if r := recover(); r != nil {
		fmt.Println("recovered from", r)
		// debug.PrintStack()
	}
}

// When a function encounters a panic, its execution is stopped, any deferred
// functions are executed and then the control returns to its caller. This
// process continues until all the functions of the current goroutine have
// returned at which point the program prints the panic message, followed
// by the stack trace and then terminates.
func fullName(firstName *string, lastName *string) {
	defer recoverFullName()
	if firstName == nil {
		panic("runtime error: first name cannot be nil")
	}
	if lastName == nil {
		panic("runtime error: last name cannot be nil")
	}
	fmt.Printf("%s %s\n", *firstName, *lastName)
	fmt.Println("returned normally from fullName")
}

func Example_panicAndRecover() {
	defer fmt.Println("deferred call in main")
	firstName := "Elon"
	fullName(&firstName, nil)
	fmt.Println("returned normally from main")
	// Output:
	// recovered from runtime error: last name cannot be nil
	// returned normally from main
	// deferred call in main
}
