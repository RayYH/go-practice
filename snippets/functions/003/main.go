package main

import (
	"fmt"
	"log"
	"runtime"
)

func main() {
	where := func() {
		_, file, line, _ := runtime.Caller(1)
		log.Printf("%s:%d\n", file, line)
	}

	where()

	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Println()
	where()
}
