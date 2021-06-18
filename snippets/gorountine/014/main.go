package main

import (
	"fmt"
)

func sendData(ch chan string) {
	ch <- "Washington"
	ch <- "Tokyo"
	ch <- "Beijing"
	ch <- "Shanghai"
	ch <- "Nanjing"
	close(ch)
}

func getData(ch chan string) {
	for {
		input, open := <-ch
		if !open {
			break
		}
		fmt.Printf("%s ", input)
	}
}

func main() {
	ch := make(chan string)
	go sendData(ch)
	getData(ch)
}
