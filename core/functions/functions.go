package main

import "fmt"

func multiPly3Nums(a int, b int, c int) int {
	// var product int = a * b * c
	// return product
	return a * b * c
}

func getX2AndX3(input int) (int, int) {
	return 2 * input, 3 * input
}

func getNamedX2AndX3(input int) (x2, x3 int) {
	x2 = input * 2
	x3 = input * 3

	return
}

// this function changes reply:
func multiply(a, b int, reply *int) {
	*reply = a * b
}

func min(s ...int) int {
	if len(s) == 0 {
		return 0
	}

	min := s[0]
	for _, v := range s {
		if v < min {
			min = v
		}
	}

	return min
}

func greet() {
	fmt.Println("Hello")
}

func connectToDB() {
	fmt.Println("ok, connected to db")
}

func disconnectFromDB() {
	fmt.Println("ok, disconnected from db")
}

func doDBOperations() {
	connectToDB()
	fmt.Println("Deferring the database disconnect.")
	defer disconnectFromDB() //function called here with defer
	fmt.Println("Doing some DB operations ...")
	fmt.Println("Oops! some crash or network error ...")
	fmt.Println("Returning from function here!")
	return //terminate the program
	// deferred function executed here just before actually returning, even if
	// there is a return or abnormal termination before
}

func trace(s string) string {
	fmt.Println("entering:", s)
	return s
}

func un(s string) {
	fmt.Println("leaving:", s)
}

func a() {
	defer un(trace("a"))
	fmt.Println("in a")
}

func b() {
	defer un(trace("b"))
	fmt.Println("in b")
	a()
}

func fibonacci(n int) (res int) {
	if n <= 1 {
		return 1
	}
	res = fibonacci(n-1) + fibonacci(n-2)
	return
}

func even(n int) bool {
	if n == 0 {
		return true
	}

	return odd(revSign(n) - 1)

}

func odd(n int) bool {
	if n == 0 {
		return false
	}

	return even(revSign(n) - 1)
}

func revSign(n int) int {
	if n < 0 {
		return -n
	}

	return n
}

func add(a, b int) int {
	return a + b
}

func callback(x, y int, f func(int, int) int) int {
	return f(x, y)
}

func getRet() (ret int) {
	defer func() {
		ret++
	}()
	return 1
}
