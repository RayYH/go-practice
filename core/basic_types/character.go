package basic_types

import "fmt"

// DisplayRuneLiterals
// \a   U+0007 alert or bell
// \b   U+0008 backspace
// \f   U+000C form feed
// \n   U+000A line feed or newline
// \r   U+000D carriage return
// \t   U+0009 horizontal tab
// \v   U+000b vertical tab
// \\   U+005c backslash
// \'   U+0027 single quote  (valid escape only within rune literals)
// \"   U+0022 double quote  (valid escape only within string literals)
func DisplayRuneLiterals() {
	fmt.Println([]byte("café"))
	fmt.Println([]rune("café"))
	runes := []rune{
		'a',
		'本',
		'\t',
		'\000',
		'\007',
		'\377',
		'\x07',
		'\xff',
		'\u12e4',
		'\U00101234',
		'\'', // rune literals containing single quote character
	}

	for index, word := range runes {
		fmt.Printf("%#U starts at byte position %d\n", word, index)
	}
}
