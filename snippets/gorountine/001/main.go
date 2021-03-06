package main

import (
	"fmt"
	"time"
)

func longWait() {
	fmt.Println("Beginning longWait()")
	time.Sleep(5 * time.Second)
	fmt.Println("End of longWait()")
}

func shortWait() {
	fmt.Println("Beginning shortWait()")
	time.Sleep(2 * time.Second)
	fmt.Println("End of shortWait()")
}

func main() {
	fmt.Println("In main()")
	// go keyword used for creating a goroutine
	go longWait()
	go shortWait()
	fmt.Println("About to sleep in main()")
	time.Sleep(10 * time.Second)
	fmt.Println("End of main()")
}
