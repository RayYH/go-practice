package main

import (
	"fmt"
	"unicode"
)

func main() {
	// \u0085 (NEL)
	// \u00A0 (NBSP)
	spaces := []rune{'\t', '\n', '\v', '\f', '\r', ' ', '\u0085', '\u00A0'}
	for _, space := range spaces {
		fmt.Println(unicode.IsSpace(space))
	}
}
