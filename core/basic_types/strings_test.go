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

func TestZeroValueOfStringType(t *testing.T) {
	var aStr string
	assert.Equal(t, aStr, "")
}

func TestStringDeclarations(t *testing.T) {
	var str string
	str = "Hello World"
	assert.Equal(t, "Hello World", str)
	assert.Equal(t, str[0], uint8('H'))
}

func TestStringTypes(t *testing.T) {
	assert.Equal(t, "\n\r", "\n\r")
	assert.Equal(t, "\\n\\r", `\n\r`)
}

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

func TestStringsLength(t *testing.T) {
	// 6 + 3 * 2 = 12
	myStr := "Hello 世界"
	assert.Equal(t, 12, len(myStr))
	assert.Equal(t, 8, utf8.RuneCountInString(myStr))
}

func TestStringIteration(t *testing.T) {
	str := "Hello,世界"
	asciiStr := "Hello World"
	t.Run("when using len method, parsed as bytes", func(t *testing.T) {
		n := len(str)
		for i := 0; i < n; i++ {
			ch := str[i]
			assert.Equal(t, fmt.Sprint(reflect.TypeOf(ch)), "uint8")
		}

		n = len(asciiStr)
		for i := 0; i < n; i++ {
			ch := asciiStr[i]
			assert.Equal(t, fmt.Sprint(reflect.TypeOf(ch)), "uint8")
		}
	})

	t.Run("as char (rune)", func(t *testing.T) {
		for _, ch := range str {
			assert.Equal(t, fmt.Sprint(reflect.TypeOf(ch)), "int32")
		}

		for _, ch := range asciiStr {
			assert.Equal(t, fmt.Sprint(reflect.TypeOf(ch)), "int32")
		}
	})

	t.Run("convert to rune slice", func(t *testing.T) {
		str2 := []rune(str)
		for i := 0; i < len(str2); i++ {
			ch := str2[i]
			assert.Equal(t, fmt.Sprint(reflect.TypeOf(ch)), "int32")
		}
	})
}

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

func TestStringsBasicOperations(t *testing.T) {
	s := "a-b-c"
	// accessing elements
	assert.Equal(t, "a", string(s[0]))
	// split string
	assert.Equal(t, []string{"a", "b", "c"}, strings.Split(s, "-"))
	// join array
	assert.Equal(t, "a-b-c", strings.Join([]string{"a", "b", "c"}, "-"))
	// trim spaces
	assert.Equal(t, "a b c", strings.TrimSpace(" a b c "))
	// concatenate strings
	assert.Equal(t, "Hello World", "Hello"+" "+"World")
}

func TestStringConversions(t *testing.T) {
	t.Run("string to int", func(t *testing.T) {
		s := "10"
		i, err := strconv.Atoi(s)
		assert.Equal(t, 10, i)
		assert.Nil(t, err)
	})

	t.Run("int to string", func(t *testing.T) {
		n := 99
		str := strconv.Itoa(n)
		assert.Equal(t, "99", str)
	})

	t.Run("any to string", func(t *testing.T) {
		var a interface{}
		a = "str"
		str2 := a.(string)
		assert.Equal(t, "str", str2)
	})
}

func ExampleDisplayStringLiterals() {
	DisplayStringLiterals()
	// Output: abc\n
	//		\n
	//"Hello, world!
	//日本語日本語日本語日本語日本語日本語
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
