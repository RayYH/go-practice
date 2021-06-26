package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("%q\n", strings.SplitN("a,b,c", ",", 2)) // ["a" "b,c"]
	z := strings.SplitN("a,b,c", ",", 0)
	fmt.Printf("%q (nil = %v)\n", z, z == nil) // [] (nil = true)
}
