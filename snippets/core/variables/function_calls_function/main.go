package main

var a string

func f2() {
	// always print global variable a
	print(a)
}

func f1() {
	// print local variable a
	a := "O"
	print(a)

	// f2 is a function, so prints global variable a
	f2()

	// f3 is a closure, so prints local variable a
	f3 := func() {
		print(a)
	}
	f3()
}

func main() {
	a = "G"
	print(a)
	f1()
}
