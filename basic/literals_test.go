package basic

import "fmt"

func ExampleDisplayIntegerLiterals() {
	numbers := []int64{
		42,                   // 十进制
		4_2,                  // 下划线会被编译器去除
		0600,                 // 0 开头代表 8 进制，6*8^2 = 6 * 64 = 384
		0_600,                // 同上
		0o600,                // 同上 0o
		0xBadFace,            // 0x 代表十六进制
		0xBad_Face,           // 同上
		0x_67_7a_2f_cc_40_c6, // 下划线会被去除
		1701411834604692310,  // 普通的十进制
		170_141183_460469_30, // 带下划线的十进制
	}

	var num int64
	for _, num = range numbers {
		fmt.Printf("%d ", num)
	}
	// Output: 42 42 384 384 384 195951310 195951310 113774485586118 1701411834604692310 17014118346046930
}

func ExampleDisplayFloatPointLiterals() {
	// with a form starting with “0x”, then hexadecimal digits, optionally including a radix point, then a “p”,
	// then a signed exponent in decimal, which is a power of two. E.g., 0x1.fp-2 is (1 + 15/16)•2^(-2) = .484375.

	numbers := []float64{
		0.,          // 0.0
		72.40,       // 72.40
		072.40,      // == 72.40
		2.71828,     // 2.71828
		1.e+0,       // 1*10^0 = 1
		6.67428e-11, // 6.67428 * 1-^-11 = 0.00000000...
		1e6,         // 1*10^6 = 1000000.0
		.25,         // 0.25
		.12345e+5,   // 12345.0
		1_5.,        // 15.0
		0.15e+0_2,   // 15.0
		0x1p-2,      // 0.25 - 十六进制的 1 乘 2 的 -2 次方 = 1 * 2^-2 = 0.25
		0x2.p10,     // 2048.0 - 2 * 2^10 = 2^11 = 2048
		0x1.Fp+0,    // 1.9375 = (1+15/16) = 1.9375
		0x.8p-0,     // 0.5 = (0+8/16) = 0.5
		0x_1FFFp-16, // 0.1249847412109375 = (1FFF) * 2^(-16) = 8191/(2^16) = 0.12498474121
		0x15e - 2,   // 0x15e - 2 (integer subtraction) = 1 * 16^2 + 5 * 16 + 14 - 2 = 256+80+14-2 = 348.0
	}

	var num float64
	for _, num = range numbers {
		fmt.Printf("%.3f ", num)
		// Output: 0.000 72.400 72.400 2.718 1.000 0.000 1000000.000 0.250 12345.000 15.000 15.000 0.250 2048.000 1.938 0.500 0.125 348.000
	}
}
