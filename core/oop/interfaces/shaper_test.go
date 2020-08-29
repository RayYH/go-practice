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
