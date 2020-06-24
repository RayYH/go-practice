package basic

import "fmt"

func ExampleInitializeSlice() {
	names := []string{"leo", "jessica", "paul"}
	checks := make([]bool, 10)
	scores := make([]int, 0, 20)
	fmt.Println(names)
	fmt.Println(checks)
	fmt.Println(scores)
	// Output: [leo jessica paul]
	// [false false false false false false false false false false]
	// []
}
