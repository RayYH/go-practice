package main

import "fmt"

func main() {
	// 在格式化字符串里，
	// %d 用于格式化整数 - %x 和 %X 用于格式化 16 进制表示的数字
	// %g 用于格式化浮点型 - %f 输出浮点数，%e 输出科学计数表示法
	fmt.Printf("%d\n", 123)      // 123
	fmt.Printf("%08d\n", 123)    // 00000123
	fmt.Printf("% 8d\n", 123)    //      123
	fmt.Printf("%-8dEnd\n", 123) // 123     End
	fmt.Printf("%x\n", 123)      // 7b
	fmt.Printf("%X\n", 123)      // 7B

	fmt.Printf("%g\n", 3.1415926)   // 3.1415926
	fmt.Printf("%f\n", 3.1415926)   // 3.141593
	fmt.Printf("%e\n", 3.1415926)   // 3.141593e+00
	fmt.Printf("%.2f\n", 3.1415926) // 3.14
	fmt.Printf("%.2g\n", 3.1415926) // 3.1
	fmt.Printf("%.2e\n", 3.1415926) // 3.14e+00
}
