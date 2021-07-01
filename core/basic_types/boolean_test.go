package basic_types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// 布尔类型的零值是 false，两个类型相同的值可以使用逻辑运算符来产生一个布尔值
func TestBooleanType(t *testing.T) {
	var aBool bool
	assert.Equal(t, false, aBool)
}
