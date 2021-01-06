package slices

import "fmt"

// 0 <= len(slice) <= cap(slice)
// declaration syntax: var identifier []type
// create an slice of an array: var slice []type = arr[start:end]
// slice is a reference

func SliceForRange() {
	seasons := []string{"Spring", "Summer", "Autumn", "Winter"}

	// element
	for ix, season := range seasons {
		fmt.Printf("Season %d is %s\n", ix, season)
	}

	// only value
	var season string
	for _, season = range seasons {
		fmt.Printf("%s\n", season)
	}

	// only index
	for ix := range seasons {
		fmt.Printf("%d\n", ix)
	}
}

func ReSlice() {
	capNum := 10
	slice := make([]int, 0, capNum)

	for i := 0; i < cap(slice); i++ {
		slice = slice[0 : i+1]
		slice[i] = i
		fmt.Printf("The length of slice[%d:%d] is %d\n", 0, i+1, len(slice))
	}

	for i := 0; i < len(slice); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice[i])
	}
}
