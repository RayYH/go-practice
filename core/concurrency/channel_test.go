package concurrency

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
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
	resultChan := make(chan int, 2)
	go sum(values[:len(values)/2], resultChan)
	go sum(values[len(values)/2:], resultChan)
	// sum1 maybe 15 or 40
	sum1, sum2 := <-resultChan, <-resultChan
	assert.Equal(t, 55, sum1+sum2)
}
