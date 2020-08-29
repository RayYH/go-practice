package identifiers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidIdentifiers(t *testing.T) {
	a := 1
	_x9 := 2
	ThisVariableIsExported := 3
	assert.Equal(t, 1, a)
	assert.Equal(t, 2, _x9)
	assert.Equal(t, 3, ThisVariableIsExported)
	// αβ := 4
	// assert.Equal(t, 4, αβ)
}
