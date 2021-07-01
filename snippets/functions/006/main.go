package main

import (
	"io"
	"log"
)

func myFunc(s string) (n int, err error) {
	defer func() {
		log.Printf("myFunc(%q) = %d, %v", s, n, err)
	}()
	return 7, io.EOF
}

func main() {
	_, _ = myFunc("Go")
}
