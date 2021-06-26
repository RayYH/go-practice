package main

import (
	"fmt"
	"strings"
)

func main() {
	// Compare this example to the ToTitle example.
	fmt.Println(strings.Title("her royal highness")) // Her Royal Highness
	fmt.Println(strings.Title("loud noises"))        // Loud Noises
	fmt.Println(strings.Title("хлеб"))               // Хлеб
}
