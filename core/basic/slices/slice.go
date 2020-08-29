package slices

import "fmt"

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

func SliceForRange() {
	var slice = make([]int, 4)
	slice[0] = 1
	slice[1] = 2
	slice[2] = 3
	slice[3] = 4
	for ix, value := range slice {
		fmt.Printf("Slice at %d is: %d\n", ix, value)
	}
}

func SliceForRange2() {
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
		fmt.Printf("The length of slice is %d\n", len(slice))
	}

	for i := 0; i < len(slice); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice[i])
	}
}

func GenerateSliceFromString() {
	s := "\u00ff\u754c"
	for i, c := range s {
		fmt.Printf("%d:%c ", i, c)
	}
}
