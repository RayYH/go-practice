package interfaces

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSquareArea(t *testing.T) {
	sq := new(Square)
	sq.side = 5.0

	var areaInterface Shaper
	areaInterface = sq
	assert.Equal(t, 25.0, areaInterface.Area())
}

func TestArea(t *testing.T) {
	r := &Square{5.0}
	q := &Rectangle{length: 5.0, width: 5.0}

	shapes := []Shaper{r, q}

	for _, shape := range shapes {
		assert.Equal(t, shape.Area(), 25.0)
	}
}

func TestTypeChecking(t *testing.T) {
	var s Shaper
	// A pointer value
	s = new(Rectangle)
	_, ok := s.(*Square)
	assert.False(t, ok)
	_, ok = s.(*Rectangle)
	assert.True(t, ok)
}

func TestGetType(t *testing.T) {
	var s1 Shaper
	var s2 Shaper
	s1 = &Rectangle{length: 1, width: 2}
	s2 = &Square{side: 1}
	assert.Equal(t, "Rec", GetType(s1))
	assert.Equal(t, "Squ", GetType(s2))
}

func ExampleClassifierCaller() {
	ClassifierCaller()
	// Output:
	// Param #0 is a int
	// Param #1 is a float64
	// Param #2 is a string
	// Param #3 is unknown
	// Param #4 is a nil
	// Param #5 is a bool
}
