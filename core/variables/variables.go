package variables

import (
	"math"
	"os"
)

// global variables
// It's visible in all files of the package to which it belongs
var globalString = "This is a string."
var (
	myName = "Ray"
	myAge  = 24
)

// We can use some built-in functions when declaring global variables
var (
	HOME = os.Getenv("HOME")
	USER = os.Getenv("USER")
)

// No initialization but only declaration
var declaredVariable float64

func init() {
	// we can initialize variables inside the init func
	declaredVariable = math.Atan(1)
}
