package variables

import (
	"math"
	"os"
)

// globalString is visible within current package.
var globalString = "This is a string."

// GlobalString is also visible to other application which imported this package.
var GlobalString = "This is also a string."

// group of global variables.
var (
	myName = "Ray"
	myAge  = 24
)

// We can use some built-in functions when declaring global variables.
var (
	HOME = os.Getenv("HOME")
	USER = os.Getenv("USER")
)

// No initialization but only declaration.
var declaredVariable float64

func init() {
	// We can initialize variables inside the init function
	declaredVariable = math.Atan(1)
}
