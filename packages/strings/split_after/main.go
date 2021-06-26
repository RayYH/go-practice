package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("%q\n", strings.SplitAfter("a,b,c", ","))      // ["a," "b," "c"]
	fmt.Printf("%q\n", strings.SplitAfter("Go Go goos", "o")) // ["Go" " Go" " go" "o" "s"]
}
