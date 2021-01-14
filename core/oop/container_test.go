package oop

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBothValueTypeAndPointerTypeWorkOnStructs(t *testing.T) {
	// value type
	var t1 IntContainer
	t1.change()
	assert.Equal(t, t1.toString(), "{1}")

	// pointer type
	t2 := new(IntContainer)
	t2.change()
	assert.Equal(t, t2.toString(), "{1}")
}

func TestInspectMixinContainer(t *testing.T) {
	mixinContainer := MixinContainer{
		BoolValue:   true,
		StringValue: "Ray Hong",
		IntValue:    1,
	}
	assert.Equal(t, "a bool value", inspectField(mixinContainer, 0))
	assert.Equal(t, "a string value", inspectField(mixinContainer, 1))
	assert.Equal(t, "an int value", inspectField(mixinContainer, 2))
}

func TestAnonymousField(t *testing.T) {
	b := BooleanContainer{}
	b.bool = true
	assert.True(t, b.bool)
}
