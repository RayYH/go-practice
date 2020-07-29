package main

// 将一个大叔组传递给函数会带来内存的大量消耗
// 有两种方法可以避免这种现象，传递数组的指针或传递数组的切片 (一般都是传递切片)

// 传递数组指针
func Sum(a *[3]float64) (sum float64) {
	for _, v := range a {
		sum += v
	}

	return
}

// 传递切片
func SliceSum(a []float64) (sum float64) {
	for _, v := range a {
		sum += v
	}

	return
}
