package basic_types

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// 1. 布尔型的值只可以是常量 true 或者 false
// 2. 布尔类型的零值是 false
// 3. 两个类型相同的值可以使用逻辑运算符来产生一个布尔值
// 4. 布尔值多用于流程控制的条件判断
func TestBooleanType(t *testing.T) {
	var aBool bool
	assert.Equal(t, false, aBool)
}
