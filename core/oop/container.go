package oop

import (
	"fmt"
	"reflect"
)

type IntContainer struct {
	thing int
}

func (b *IntContainer) change()         { b.thing = 1 }
func (b IntContainer) toString() string { return fmt.Sprint(b) }

type BooleanContainer struct {
	bool // anonymous field
}

type MixinContainer struct {
	BoolValue   bool   `description:"a bool value"`
	StringValue string `description:"a string value"`
	IntValue    int    `description:"an int value"`
}

func inspectField(mc MixinContainer, i int) string {
	r := reflect.TypeOf(mc)
	field := r.Field(i)
	return fmt.Sprintf("%v", field.Tag.Get("description"))
}
