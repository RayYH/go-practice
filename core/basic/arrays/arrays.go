package arrays

// Passing a large array to a function will cause a lot of memory consumption,
// because the array is a value type, and a copy operation will occur when passed as a parameter.
// There are two ways to avoid this, 1) pass the pointer of the array; 2) pass the slice of the array.

// Pass a pointer to an array of [3]float64 type
func Sum(a *[3]float64) (sum float64) {
	for _, v := range a {
		sum += v
	}

	return
}

// Pass a slice of []float64
func SliceSum(a []float64) (sum float64) {
	for _, v := range a {
		sum += v
	}

	return
}

// Accepting an array as a parameter does not change the value in the original array
func TryToModify(arr [5]int) {
	arr[0] = 0
}
