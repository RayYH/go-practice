package basic_types

import "fmt"

func DisplayStringLiterals() {
	strings := []string{
		`abc`,
		`\n
		\n`,
		"\n",
		"\"",
		"Hello, world!\n",
		"日本語",
		"\u65e5本\U00008a9e",
		`日本語`,
		"\u65e5\u672c\u8a9e",
		"\U000065e5\U0000672c\U00008a9e",
		"\xe6\x97\xa5\xe6\x9c\xac\xe8\xaa\x9e",
	}

	var value string
	for _, value = range strings {
		fmt.Printf("%s", value)
	}
}
