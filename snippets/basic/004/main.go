package main

import (
	"fmt"
	"os"
	"runtime"
)

func main() {
	var goos = runtime.GOOS
	fmt.Printf("The operating system is %s\n", goos)
	path := os.Getenv("PATH")
	fmt.Printf("The path is %s\n", path)
}
