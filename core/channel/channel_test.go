package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReadWriteChannel(t *testing.T) {
	var readWriteChannel chan int
	assert.Equal(t, "chan int", fmt.Sprintf("%T", readWriteChannel))
	assert.Nil(t, readWriteChannel)
}

func TestReadOnlyChannel(t *testing.T) {
	var readOnlyChannel <-chan int
	assert.Equal(t, "<-chan int", fmt.Sprintf("%T", readOnlyChannel))
	assert.Nil(t, readOnlyChannel)
}

func TestWriteOnlyChannel(t *testing.T) {
	var writeOnlyChannel chan<- int
	assert.Equal(t, "chan<- int", fmt.Sprintf("%T", writeOnlyChannel))
	assert.Nil(t, writeOnlyChannel)
}

func TestChannel(t *testing.T) {
	var ch1 chan int
	ch2 := make(chan int)
	assert.Equal(t, "chan int", fmt.Sprintf("%T", ch1))
	assert.Equal(t, "chan int", fmt.Sprintf("%T", ch2))
	assert.NotNil(t, &ch1)
	assert.NotNil(t, &ch2)
	assert.Nil(t, ch1)
	assert.NotNil(t, ch2)
}

func TestSum(t *testing.T) {
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	resultChan := make(chan int, 2)
	go sum(values[:len(values)/2], resultChan)
	go sum(values[len(values)/2:], resultChan)
	// sum1 可能是 15 也可能是 40
	sum1, sum2 := <-resultChan, <-resultChan
	assert.Equal(t, 55, sum1+sum2)
}
