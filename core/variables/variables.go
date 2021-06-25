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
	_ = os.Getenv("HOME")
	_ = os.Getenv("USER")
)

// No initialization but only declaration.
var declaredVariable float64

// init 函数不能被人为调用，该函数在每个包完成初始化后自动执行，并且执行优先级比 `main` 函数高
func init() {
	// We can initialize variables inside the init function
	declaredVariable = math.Atan(1)
}
