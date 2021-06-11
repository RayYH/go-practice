package types

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFloatingPointNumbersDivision(t *testing.T) {
	a := 7.0
	b := 3.0
	assert.Equal(t, a/b, 2.3333333333333335)
}
