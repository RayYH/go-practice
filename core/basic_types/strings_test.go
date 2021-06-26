package basic_types

import (
	"bytes"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"testing"
	"unicode/utf8"

	"github.com/stretchr/testify/assert"
)

// 字符串是 UTF-8 字符的一个序列
// 当字符为 ASCII 码时则占用 1 个字节，其它字符根据需要占用 2-4 个字节
// Go 中字符串的零值是空字符串
func TestZeroValueOfStringType(t *testing.T) {
	// see strings/strings_test for more information
	var aStr string
	assert.Equal(t, aStr, "")
}

// Go 中的字符串可以通过下标访问，但是不能通过下标修改，这同其他大多数语言一致 (PHPer 沉默了)
func TestStringDeclarations(t *testing.T) {
	var str string
	str = "Hello World"
	assert.Equal(t, "Hello World", str)
	assert.Equal(t, str[0], uint8('H'))
}

// byte/rune 切片可以直接转换为字符串
func TestSlicesCanBeConvertedToStrings(t *testing.T) {
	// string literal
	var givenString = "Geeks"
	// Creating and initializing a slices of byte
	mySlice1 := []byte{0x47, 0x65, 0x65, 0x6b, 0x73}
	// Creating a string from the slices
	myString1 := string(mySlice1)
	// rune
	mySlice2 := []rune{0x0047, 0x0065, 0x0065, 0x006b, 0x0073}
	// creating a string from the slice
	myString2 := string(mySlice2)

	assert.Equal(t, myString1, givenString)
	assert.Equal(t, myString2, givenString)
}

// len 方法返回字符串长度，包含中文字符时，需要用 utf8.RuneCountInString
func TestStringsLength(t *testing.T) {
	// 6 + 3 * 2 = 12
	myStr := "Hello 世界"
	assert.Equal(t, 12, len(myStr))
	assert.Equal(t, 8, utf8.RuneCountInString(myStr))
}

// 由于字符串可以通过下标访问，因此字符串可以进行遍历 (迭代)
func TestStringIteration(t *testing.T) {
	str := "Hello,世界"
	// 通过 len 作为上限索引来循环，是按 ASCII 字符解析的
	t.Run("as byte", func(t *testing.T) {
		n := len(str)
		for i := 0; i < n; i++ {
			ch := str[i]
			assert.Equal(t, fmt.Sprint(reflect.TypeOf(ch)), "uint8")
		}
	})

	// 使用 range 来进行遍历时，默认是按 rune (int32) 来解析的
	t.Run("as char (rune)", func(t *testing.T) {
		for _, ch := range str {
			assert.Equal(t, fmt.Sprint(reflect.TypeOf(ch)), "int32")
		}
	})

	// 强制转换为 rune slice 然后使用索引遍历可以达到和 range 一样的效果
	t.Run("convert to rune slice", func(t *testing.T) {
		str2 := []rune(str)
		for i := 0; i < len(str2); i++ {
			ch := str2[i]
			assert.Equal(t, fmt.Sprint(reflect.TypeOf(ch)), "int32")
		}
	})
}

// Go 支持使用 `` 来表示多行字符串 (JavaScript 风格)
func TestMultiLineStrings(t *testing.T) {
	t.Run("empty multiple line string", func(t *testing.T) {
		var rawStr1 = ``
		assert.Equal(t, "", rawStr1)
	})

	t.Run("non-empty multiple line string", func(t *testing.T) {
		var multilineStr = `
    1234567890
    ~!@#$%^&*()_+|
    abcABC
    `
		assert.Equal(t, "\n    1234567890\n    ~!@#$%^&*()_+|\n    abcABC\n    ", multilineStr)
	})

	t.Run("xml", func(t *testing.T) {
		var xmlStr1 = `
<xml>
    <node></node>
</xml>
`
		assert.Equal(t, "\n<xml>\n    <node></node>\n</xml>\n", xmlStr1)
	})

	t.Run("sql", func(t *testing.T) {
		var sql = `
SELECT *
FROM table
`
		assert.Equal(t, "\nSELECT *\nFROM table\n", sql)
	})

	t.Run("json", func(t *testing.T) {
		var json = `
{
    "key": "value"
}
`
		assert.Equal(t, "\n{\n    \"key\": \"value\"\n}\n", json)
	})
}

// Go 字符串基础操作
func TestStringsBasicOperations(t *testing.T) {
	s := "a-b-c"
	// accessing elements - 访问
	assert.Equal(t, "a", string(s[0]))
	// split string - 切割
	assert.Equal(t, []string{"a", "b", "c"}, strings.Split(s, "-"))
	// join array - 合并
	assert.Equal(t, "a-b-c", strings.Join([]string{"a", "b", "c"}, "-"))
	// trim spaces - 裁剪
	assert.Equal(t, "a b c", strings.TrimSpace(" a b c "))
	// concatenate strings - 连接
	assert.Equal(t, "Hello World", "Hello"+" "+"World")
}

// 字符串转换，主要是与数值类型，用到 strconv 包
func TestStringConversions(t *testing.T) {
	// string to int
	s := "10"
	i, err := strconv.Atoi(s)
	assert.Equal(t, 10, i)
	assert.Nil(t, err)

	// int to string
	n := 99
	str := strconv.Itoa(n)
	assert.Equal(t, "99", str)

	// a can be any type
	var a interface{}
	a = "str"
	str2 := a.(string)
	assert.Equal(t, "str", str2)
}

// using +=: 10.57 ns/op
func BenchmarkStringAdd(b *testing.B) {
	b.ResetTimer()
	s := ""
	for i := 0; i < b.N; i++ {
		s += "s"
		s = ""
	}
	b.StopTimer()
}

// using Sprintf: 102.9 ns/op
func BenchmarkSprintf(b *testing.B) {
	b.ResetTimer()
	s := "s"
	for i := 0; i < b.N; i++ {
		s = fmt.Sprintf("%v", s)
	}
	b.StopTimer()
}

// using strings.Builder - 3.350 ns/op
func BenchmarkStringsBuilder(b *testing.B) {
	b.ResetTimer()
	var stringsBuilder strings.Builder
	s := "s"
	for i := 0; i < b.N; i++ {
		stringsBuilder.WriteString(s)
	}
	_ = stringsBuilder.String()
	b.StopTimer()
}

// using bytes.Buffer - 11.67 ns/op
func BenchmarkBytesBuffer(b *testing.B) {
	b.ResetTimer()
	s := "s"
	var bytesBuffer bytes.Buffer
	for i := 0; i < b.N; i++ {
		bytesBuffer.WriteString(s)
	}
	_ = bytesBuffer.String()
	b.StopTimer()
}
