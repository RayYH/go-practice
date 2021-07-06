package main

import "fmt"

func main() {
	i := 42
	j := 3.14
	k := 0.867 + 0.5i

	fmt.Printf("%T %T %T", i, j, k)
}
