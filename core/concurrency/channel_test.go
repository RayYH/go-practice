package concurrency

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

// Channels are a typed conduit through which you can
// send and receive values with the channel operator: <-.
func TestChannelTypes(t *testing.T) {
	var readWriteChannel chan int
	assert.Equal(t, "chan int", fmt.Sprintf("%T", readWriteChannel))
	assert.Nil(t, readWriteChannel)

	var readOnlyChannel <-chan int
	assert.Equal(t, "<-chan int", fmt.Sprintf("%T", readOnlyChannel))
	assert.Nil(t, readOnlyChannel)

	var writeOnlyChannel chan<- int
	assert.Equal(t, "chan<- int", fmt.Sprintf("%T", writeOnlyChannel))
	assert.Nil(t, writeOnlyChannel)
}

func TestSum(t *testing.T) {
	sum := func(values []int, resultChan chan int) {
		total := 0
		for _, v := range values {
			total += v
		}

		resultChan <- total
	}

	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// Channels can be buffered.
	// Provide the buffer length as the second argument to make to initialize a buffered channel.
	resultChan := make(chan int, 2)
	go sum(values[:len(values)/2], resultChan)
	go sum(values[len(values)/2:], resultChan)
	// sum1 maybe 15 or 40
	sum1, sum2 := <-resultChan, <-resultChan
	assert.Equal(t, 55, sum1+sum2)
}

func TestCloseAChannel(t *testing.T) {
	fibonacci := func(n int, c chan int) {
		x, y := 0, 1
		for i := 0; i < n; i++ {
			c <- x
			x, y = y, x+y
		}
		close(c)
	}
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	arr := [10]int{}
	index := 0
	time.Sleep(1500 * time.Millisecond)
	for i := range c {
		arr[index] = i
		index += 1
	}
	assert.Equal(t, [10]int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}, arr)
}

func TestGoroutineWaitOnMultipleCommunicationOperations(t *testing.T) {
	// A select blocks until one of its cases can run, then it executes that case.
	// It chooses one at random if multiple are ready.
	fibonacci := func(c, quit chan int) {
		x, y := 0, 1
		for {
			select {
			case c <- x:
				x, y = y, x+y
			case <-quit:
				return
			}
		}
	}

	c := make(chan int)
	quit := make(chan int)
	go func(t *testing.T) {
		arr := [10]int{0, 1, 1, 2, 3, 5, 8, 13, 21, 34}
		for i := 0; i < 10; i++ {
			assert.Equal(t, arr[i], <-c)
		}
		quit <- 0
	}(t)
	fibonacci(c, quit)
}
