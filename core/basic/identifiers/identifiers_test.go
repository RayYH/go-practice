package identifiers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// a name begins with a letter or an underscore
// Case matters: heapSort and HeapSort are different names
func TestValidIdentifiers(t *testing.T) {
	a := 0
	A := 1
	_x9 := 2
	ThisVariableIsExported := 3
	assert.Equal(t, 0, a)
	assert.Equal(t, 1, A)
	assert.Equal(t, 2, _x9)
	assert.Equal(t, 3, ThisVariableIsExported)
	// αβ := 4
	// assert.Equal(t, 4, αβ)
}
