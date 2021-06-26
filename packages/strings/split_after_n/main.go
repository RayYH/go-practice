package main

import (
	"fmt"
	"strings"
)

func main() {
	// n > 0: at most n substrings; the last substring will be the unsplit remainder.
	// n == 0: the result is nil (zero substrings)
	// n < 0: all substrings
	fmt.Printf("%q\n", strings.SplitAfterN("a,b,c", ",", 2)) // ["a," "b,c"]
}
