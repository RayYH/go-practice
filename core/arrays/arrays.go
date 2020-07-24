package main

// 把一个大数组传递给函数会消耗很多内存。有两种方法可以避免这种现象，传递数组的指针，使用数组的切片
// 一般都是使用切片而不是数组指针
func Sum(a *[3]float64) (sum float64) {
	for _, v := range a {
		sum += v
	}

	return
}
