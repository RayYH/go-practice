package main

import (
	"fmt"
	"strings"
)

func main() {
	var s = "¡¡¡Hello, Gophers!!!"
	s = strings.TrimSuffix(s, ", Gophers!!!")
	fmt.Println(s)
	s = strings.TrimSuffix(s, ", Marmots!!!")
	fmt.Println(s)
}
