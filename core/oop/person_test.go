package oop

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPersonStruct(t *testing.T) {
	// struct as a value type
	var person1 Person
	person1.firstName = "Ray"
	person1.lastName = "Hong"
	upperPerson(&person1)
	assert.Equal(t, person1.firstName, "RAY")
	assert.Equal(t, person1.lastName, "HONG")

	// struct as a pointer
	person2 := new(Person)
	person2.firstName = "Ray"
	person2.lastName = "Hong"
	upperPerson(person2)
	assert.Equal(t, person2.firstName, "RAY")
	assert.Equal(t, person2.lastName, "HONG")

	// struct as a literal
	person3 := &Person{firstName: "Ray", lastName: "Hong"}
	upperPerson(person3)
	assert.Equal(t, person3.firstName, "RAY")
	assert.Equal(t, person3.lastName, "HONG")
}

func TestPersonString(t *testing.T) {
	person := &Person{firstName: "Ray", lastName: "Hong"}
	assert.Equal(t, fmt.Sprint(person), "Ray Hong")
}
