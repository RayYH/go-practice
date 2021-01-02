package variables

import (
	"math"
	"os"
)

// global variables - it's visible in all files of the package to which it belongs.
var globalString = "This is a string."

var (
	myName = "Ray"
	myAge  = 24
)

// use some built-in functions
var (
	HOME = os.Getenv("HOME")
	USER = os.Getenv("USER")
)

var declaredVariable float64

func init() {
	// we can initialize variables here
	declaredVariable = math.Atan(1)
}
