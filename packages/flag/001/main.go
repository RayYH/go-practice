// These examples demonstrate more intricate uses of the flag package.
package main

import (
	"flag"
	"fmt"
)

// Example 1: A single string flag called "species" with default value "gopher".
var species = flag.String("species", "gopher", "the species we are studying")

func main() {
	// All the interesting pieces are with the variables declared above, but
	// to enable the flag package to see the flags defined there, one must
	// execute, typically at the start of main (not init!):
	//	flag.Parse()
	// We don't run it here because this is not a main function and
	// the testing suite has already parsed the flags.
	flag.Parse()
	fmt.Println("species: ", *species)
}
