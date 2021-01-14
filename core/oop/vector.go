package oop

// alias []int IntVector
type IntVector []int

func (v IntVector) Sum() (s int) {
	for _, value := range v {
		s += value
	}
	return
}
