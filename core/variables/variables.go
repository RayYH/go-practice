package variables

import (
	"math"
	"os"
	"runtime"
)

// global variables
var globalString = "This is a string."
var GlobalString = "This is also a string."

// only declaration
var emptyGlobalVar string
var declaredVariable float64

var (
	myName = "Ray"
	myAge  = 24
)

// use built-in functions
var (
	_ = runtime.Version()
	_ = os.Getenv("HOME")
	_ = os.Getenv("USER")
)

func init() {
	declaredVariable = math.Atan(1)
}
