package oop

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStructFactory(t *testing.T) {
	f := NewFile(1, "filename")
	assert.Equal(t, f.fd, 1)
	assert.Equal(t, f.name, "filename")
}
