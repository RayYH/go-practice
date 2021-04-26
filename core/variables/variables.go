package variables

import (
	"math"
	"os"
)

// Global variables have a global scope, they are visible from
// anywhere within the application (or package).
// globalString is visible within current package.
// GlobalString is also visible to other application which imported this package.
var globalString = "This is a string."
var GlobalString = "This is also a string."
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
