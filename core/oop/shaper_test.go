package oop

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSquareArea(t *testing.T) {
	sq := new(Square)
	sq.side = 5.0

	var box Shaper
	box = sq
	assert.Equal(t, 25.0, box.Area())
}

func TestArea(t *testing.T) {
	s := &Square{5.0}
	r := &Rectangle{length: 5.0, width: 5.0}

	shapes := []Shaper{s, r}

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
