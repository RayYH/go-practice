package main

var a string

func f2() {
	print(a)
}

func f1() {
	a := "O"
	print(a)
	f2()
}

func main() {
	a = "G"
	print(a)
	f1()
}
