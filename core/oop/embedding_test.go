package oop

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAccessInnerAnonymousFieldsDirectly(t *testing.T) {
	s := new(outerStruct)
	s.intValue = 1
	assert.Equal(t, s.intValue, 1)
}

func TestAccessInnerFieldsViaChainingInvocation(t *testing.T) {
	s := new(outerStruct)
	s.bs.boolValue = false
	assert.Equal(t, s.bs.boolValue, false)
}

func TestNameOverridden(t *testing.T) {
	f := floatStruct{floatValue: 1.1}
	o := &outerStruct{
		floatStruct: f,
		floatValue:  2.2,
	}
	assert.Equal(t, o.floatValue, 2.2)
	assert.Equal(t, f.floatValue, 1.1)
}
