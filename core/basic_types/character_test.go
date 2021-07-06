package basic_types

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBytesPresentation(t *testing.T) {
	var chA byte = 65
	var chB byte = '\x41'
	var ch1 int = '\u0041'
	var ch2 int = '\u03B2'
	var ch3 int = '\U00101234'

	// %c - character
	// %d/%v - integer
	// %X - UTF-8 bytes
	// %U - UTF-8 code point
	assert.Equal(t, "65 65", fmt.Sprint(chA, chB)) // 65, 65
	assert.Equal(t, "65 946 1053236", fmt.Sprint(ch1, ch2, ch3))
	assert.Equal(t, fmt.Sprintf("%d - %d - %d", ch1, ch2, ch3), "65 - 946 - 1053236")
	assert.Equal(t, fmt.Sprintf("%c - %c - %c", ch1, ch2, ch3), "A - β - \U00101234")
	assert.Equal(t, fmt.Sprintf("%X - %X - %X", ch1, ch2, ch3), "41 - 3B2 - 101234")
	assert.Equal(t, fmt.Sprintf("%U - %U - %U", ch1, ch2, ch3), "U+0041 - U+03B2 - U+101234")
}

func ExampleDisplayRuneLiterals() {
	DisplayRuneLiterals()
	// Output: [99 97 102 195 169]
	// [99 97 102 233]
	// U+0061 'a' starts at byte position 0
	// U+672C '本' starts at byte position 1
	// U+0009 starts at byte position 2
	// U+0000 starts at byte position 3
	// U+0007 starts at byte position 4
	// U+00FF 'ÿ' starts at byte position 5
	// U+0007 starts at byte position 6
	// U+00FF 'ÿ' starts at byte position 7
	// U+12E4 'ዤ' starts at byte position 8
	// U+101234 starts at byte position 9
	// U+0027 ''' starts at byte position 10
}
