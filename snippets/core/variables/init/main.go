package main

import (
	"fmt"
	// https://golang.org/cmd/go/#hdr-Relative_import_paths
	// Second, if you are compiling a Go program not in a work space, you can use a
	// relative path in an import statement in that program to refer to nearby code also not in a work space.
	"github.com/rayyh/go-practice/snippets/core/variables/init/trans"
)

var twoPi = 2 * trans.Pi

func main() {
	fmt.Printf("2*Pi = %g\n", twoPi) // 2*Pi = 6.283185307179586
}
