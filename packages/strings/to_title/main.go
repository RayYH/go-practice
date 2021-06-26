package main

import (
	"fmt"
	"strings"
)

func main() {
	// Compare this example to the Title example.
	fmt.Println(strings.ToTitle("her royal highness")) // HER ROYAL HIGHNESS
	fmt.Println(strings.ToTitle("loud noises"))        // LOUD NOISES
	fmt.Println(strings.ToTitle("хлеб"))               // ХЛЕБ
}
