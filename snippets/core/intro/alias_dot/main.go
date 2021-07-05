package main

import (
	// If an explicit period (.) appears instead of a name, all the package's exported identifiers will be declared in
	// the current file's file block and can be accessed without a qualifier.
	. "fmt"
	. "math"
)

func main() {
	Printf("Abs(-1) = %.2f", Abs(-1.23))
}
