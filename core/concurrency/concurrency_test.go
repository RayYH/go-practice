package concurrency

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

// A goroutine is a lightweight thread managed by the Go runtime.
// `go f(x, y, z)` starts a new goroutine running `f(x, y, z)`
// The evaluation of f, x, y, and z happens in the current goroutine
// and the execution of f happens in the new goroutine.

func TestSimple(t *testing.T) {
	modifyArray := func(arr *[1]int) {
		time.Sleep(300 * time.Millisecond)
		arr[0] = 1
	}

	array := [1]int{0}
	go modifyArray(&array)
	assert.Equal(t, array, [1]int{0})
	time.Sleep(1500 * time.Millisecond)
	assert.Equal(t, array, [1]int{1})
}
