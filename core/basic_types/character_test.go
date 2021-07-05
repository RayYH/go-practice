package basic_types

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// byte 类型是 uint8 的别名，适用于传统的 ASCII 编码
// rune 类型是 int32 的别名，适用于 UTF-8 编码
func TestBytesPresentation(t *testing.T) {
	// byte 是 uint8 的别名，因此 byte 只能定义 ASCII 字符
	var chA byte = 65
	var chB byte = '\x41'
	// 多字节字符我们可以用 int 类型变量存储
	// 在书写 Unicode 字符时，需要在 16 进制数之前加上前缀 \u 或者 \U
	var ch1 int = '\u0041'
	var ch2 int = '\u03B2'
	var ch3 int = '\U00101234'

	// 对于同一个字符型变量，不同的解析方式 (规则) 会有不同的呈现效果

	// %c - character
	// %d/%v - integer
	// %X - UTF-8 bytes
	// %U - UTF-8 core point
	assert.Equal(t, "65 65", fmt.Sprint(chA, chB)) // 65, 65
	assert.Equal(t, "65 946 1053236", fmt.Sprint(ch1, ch2, ch3))
	assert.Equal(t, fmt.Sprintf("%d - %d - %d", ch1, ch2, ch3), "65 - 946 - 1053236")
	assert.Equal(t, fmt.Sprintf("%c - %c - %c", ch1, ch2, ch3), "A - β - \U00101234")
	assert.Equal(t, fmt.Sprintf("%X - %X - %X", ch1, ch2, ch3), "41 - 3B2 - 101234")
	assert.Equal(t, fmt.Sprintf("%U - %U - %U", ch1, ch2, ch3), "U+0041 - U+03B2 - U+101234")
}
