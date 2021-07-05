package main

var a = "G"

func printA() {
	// global a
	print(a)
}

func modifyAndPrintA() {
	// re-define a local variable a
	a := "O"
	print(a)
}

func main() {
	printA()          // "G"
	modifyAndPrintA() // "O"
	printA()          // "G"
}
