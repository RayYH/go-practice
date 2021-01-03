package functions

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestMultiPly3Nums(t *testing.T) {
	assert.Equal(t, 60, multiPly3Nums(2, 5, 6))
}

func TestGet2XAnd3X(t *testing.T) {
	x2, x3 := getX2AndX3(2)
	assert.Equal(t, 4, x2)
	assert.Equal(t, 6, x3)
}

func TestGetNamed2XAnd3X(t *testing.T) {
	x2, x3 := getNamedX2AndX3(2)
	assert.Equal(t, 4, x2)
	assert.Equal(t, 6, x3)
}

func TestMultiply(t *testing.T) {
	n := 0
	reply := &n
	multiply(3, 4, reply)
	assert.Equal(t, 12, *reply)
}

func TestMin(t *testing.T) {
	assert.Equal(t, 1, min(1, 2, 3, 4, 5))
}

func ExampleGreet() {
	defer Greet()
	fmt.Println("DEFER")
	// Output:
	// DEFER
	// Hello
}

func ExampleDoDBOperations() {
	DoDBOperations()
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
	// DEFER 3
	// DEFER 2
	// DEFER 1
}

func TestFibonacci(t *testing.T) {
	result := 0
	nums := []int{1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89}
	for i := 0; i <= 10; i++ {
		result = fibonacci(i)
		assert.Equal(t, nums[i], result)
	}
}

func TestEvenAndOdd(t *testing.T) {
	assert.True(t, odd(17))
	assert.True(t, even(18))
}

func TestCallback(t *testing.T) {
	assert.Equal(t, 3, callback(1, 2, add))
}

func ExampleClosure() {
	Closure()
	// Output:
	// 0123
}

func TestGetRet(t *testing.T) {
	assert.Equal(t, 2, getRet())
}

func TestAdder(t *testing.T) {
	plusTwo := AddTwo()
	assert.Equal(t, 5, plusTwo(3))
	plus := Adder(4)
	assert.Equal(t, 9, plus(5))
}

func TestSequentialAdder(t *testing.T) {
	var f = SequentialAdder()
	assert.Equal(t, 1, f(1))
	assert.Equal(t, 21, f(20))
	assert.Equal(t, 321, f(300))
}

func TestMakeAddSuffix(t *testing.T) {
	addBmp := MakeAddSuffix(".bmp")
	addJpeg := MakeAddSuffix(".jpeg")
	assert.Equal(t, "file.bmp", addBmp("file"))
	assert.Equal(t, "file.jpeg", addJpeg("file"))
}

func TestCalculateFunctionExecTime(t *testing.T) {
	start := time.Now()
	for i := 0; i < 1000; i++ {

	}
	end := time.Now()
	delta := end.Sub(start)
	assert.Greater(t, delta.Nanoseconds(), int64(0))
}

func ExampleMyPrintf() {
	var v1 = 1 // inferred int
	var v2 int64 = 234
	var v3 = "Hello" // inferred string
	var v4 float32 = 1.234
	MyPrintf(v1, v2, v3, v4)
	// Output:
	// 1 is an int value.
	// 234 is an int64 value.
	// Hello is a string value.
	// 1.234 is an unknown type.
}
