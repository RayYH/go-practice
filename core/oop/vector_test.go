package oop

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntVectorSum(t *testing.T) {
	assert.Equal(t, IntVector{1, 2, 3}.Sum(), 6)
}
