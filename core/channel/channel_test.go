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
