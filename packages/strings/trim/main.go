package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Trim("¡¡¡Hello, Gophers!!!", "!¡"))
	fmt.Println(strings.Trim("   Hello, Gophers    ", ""))
	fmt.Println(strings.Trim("\t      Hello, Gophers    ", " "))
	fmt.Println(strings.Trim("\t      Hello, Gophers    ", " \t"))
}
