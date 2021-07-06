package basic_types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBooleanType(t *testing.T) {
	var aBool bool
	assert.Equal(t, false, aBool)
}
