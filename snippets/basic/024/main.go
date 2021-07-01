package main

import "fmt"

func main() {
	symbol := "*"
	width, height := 20, 10
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			fmt.Print(symbol)
		}
		fmt.Println()
	}
}
