package arrays

const (
	WIDTH  = 192
	HEIGHT = 108
)

type pixel int

var screen [WIDTH][HEIGHT]pixel

// Pass a pointer to an array of [3]float64 type
func Sum(arr *[3]float64) (sum float64) {
	for _, v := range arr {
		sum += v
	}

	return
}

// Pass a slice of []float64
func SliceSum(arr []float64) (sum float64) {
	for _, v := range arr {
		sum += v
	}

	return
}
