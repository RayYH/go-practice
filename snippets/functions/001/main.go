package main

import "fmt"

func main() {
	j := 1
	a := func() {
		// is is passed by value, arr is passed by reference
		i := 0
		slice := []int{1, 2, 3}
		defer fmt.Println("[outside defer func] i =", i)
		defer fmt.Println("[outside defer func] slice =", slice)
		defer func() {
			defer fmt.Println("[inside defer func] i =", i)
			defer fmt.Println("[inside defer func] slice =", slice)
			j += i
			j += slice[0]
			slice[0] = 100
		}()
		i++
		return
	}
	a()
	fmt.Println("[finally] j =", j)
}
