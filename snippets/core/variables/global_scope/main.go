package main

var a = "G"

func n() {
	print(a)
}

func m() {
	// modify global variable a
	a = "O"
	print(a)
}

func main() {
	n() // "G"
	m() // "O"
	n() // "O"
}
