package main

// 切片是对数组一个连续片段的引用，该数组我们称之为相关数组，通常是匿名的，所以切片是一个引用类型
// 切片是一个长度可变的数组：0 <= len(s) <= cap(s)
// 声明切片的格式是： var identifier []type 无需声明长度
// 一个切片在未初始化之前默认为 nil，长度为 0
// 切片的初始化格式是：var slice []type = arr[start:end]

// len() -> 获取切片长度
// cap() -> 获取切片容量

// 因为切片是引用，所以它们不需要使用额外的内存并且比使用数组更有效率

func Compare(a, b []byte) int {
	for i := 0; i < len(a) && i < len(b); i++ {
		switch {
		case a[i] > b[i]:
			return 1
		case a[i] < b[i]:
			return -1
		}
	}
	// 数组的长度可能不同
	switch {
	case len(a) < len(b):
		return -1
	case len(a) > len(b):
		return 1
	}
	return 0 // 数组相等
}
