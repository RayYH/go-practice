package main

import (
	"fmt"
	"strings"
)

func main() {
	var s = "¡¡¡Hello, Gophers!!!"
	s = strings.TrimPrefix(s, "¡¡¡Hello, ")
	fmt.Println(s)
	s = strings.TrimPrefix(s, "¡¡¡Howdy, ")
	fmt.Println(s)
}
