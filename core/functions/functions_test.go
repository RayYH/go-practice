package main

import (
	"fmt"
	"testing"
)

func TestMultiPly3Nums(t *testing.T) {
	if multiPly3Nums(2, 5, 6) != 60 {
		t.Error("2 * 5 * 6 != 60")
	}
}

func TestGet2XAnd3X(t *testing.T) {
	if x2, x3 := getX2AndX3(2); x2 != 4 || x3 != 6 {
		t.Error("2 * 2 != 4 OR 2 * != 6")
	}
}

func TestGetNamed2XAnd3X(t *testing.T) {
	if x2, x3 := getNamedX2AndX3(2); x2 != 4 || x3 != 6 {
		t.Error("2 * 2 != 4 OR 2 * != 6")
	}
}

func TestMultiply(t *testing.T) {
	n := 0
	reply := &n
	if multiply(3, 4, reply); *reply != 12 {
		t.Error("3 * 4 != 12")
	}
}

func TestMin(t *testing.T) {
	if min(1, 2, 3, 4, 5) != 1 {
		t.Error("1 is the smallest value in 1 2 3 4 5")
	}
}

func ExampleDefer() {
	defer greet()
	fmt.Println("DEFER")
	// Output:
	// DEFER
	// Hello
}

func ExampleDoDBOperations() {
	doDBOperations()
	// Output:
	// ok, connected to db
	// Deferring the database disconnect.
	// Doing some DB operations ...
	// Oops! some crash or network error ...
	// Returning from function here!
	// ok, disconnected from db
}

func ExampleB() {
	b()
	// Output:
	// entering: b
	// in b
	// entering: a
	// in a
	// leaving: a
	// leaving: b
}

func ExampleFibonacci() {
	result := 0
	for i := 0; i <= 10; i++ {
		result = fibonacci(i)
		fmt.Printf("fibonacci(%d) is: %d\n", i, result)
	}
	// Output:
	// fibonacci(0) is: 1
	// fibonacci(1) is: 1
	// fibonacci(2) is: 2
	// fibonacci(3) is: 3
	// fibonacci(4) is: 5
	// fibonacci(5) is: 8
	// fibonacci(6) is: 13
	// fibonacci(7) is: 21
	// fibonacci(8) is: 34
	// fibonacci(9) is: 55
	// fibonacci(10) is: 89
}

func TestEvenAndOdd(t *testing.T) {
	if odd(18) {
		t.Error("18 is odd")
	}

	if even(19) {
		t.Error("19 is even")
	}
}

func TestCallback(t *testing.T) {
	if callback(1, 2, add) != 3 {
		t.Error("1+2 != 3")
	}
}

func ExampleClosure() {
	g := func(i int) {
		fmt.Printf("%d", i)
	}
	for i := 0; i < 4; i++ {
		g(i)
	}
	// Output:
	// 0123
}

func TestGetRet(t *testing.T) {
	if getRet() != 2 {
		t.Error("error")
	}
}

