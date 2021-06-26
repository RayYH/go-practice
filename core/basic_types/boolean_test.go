package basic_types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// 布尔类型的零值 (zero value) 是 false
// 布尔型的常量和变量可以通过逻辑运算符来产生一个布尔值
func TestBooleanType(t *testing.T) {
	var aBool bool
	assert.Equal(t, false, aBool)
}
