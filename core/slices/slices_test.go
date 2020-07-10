package main

import "fmt"

func ExampleInitializeSlice() {
	names := []string{"leo", "jessica", "paul"}
	checks := make([]bool, 10)
	scores := make([]int, 2, 20)
	numbers := new([20]int)[0:2]
	fmt.Println(names)
	fmt.Println(checks)
	fmt.Println(scores)
	fmt.Println(numbers)
	// Output: [leo jessica paul]
	// [false false false false false false false false false false]
	// [0 0]
	// [0 0]
}
