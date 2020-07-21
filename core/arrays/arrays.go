package main

func Sum(a *[3]float64) (sum float64) {
	for _, v := range a {
		sum += v
	}

	return
}
