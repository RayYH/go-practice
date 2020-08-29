package literals

import "fmt"

func DisplayIntegerLiterals() {
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
}

func DisplayFloatPointLiterals() {
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
	}
}

func DisplayImaginaryLiterals() {
	complexes := []complex128{
		0i,
		123i,   // == 123i for backward-compatibility
		0o123i, // == 0o123 * 1i == 83i
		0xabci, // == 0xabc * 1i == 2748i
		0.i,
		2.71828i,
		1.e+0i,
		6.67428e-11i,
		1e6i,
		.25i,
		.12345e+5i,
		0x1p-2i, // == 0x1p-2 * 1i == 0.25i
		100 + 10i,
	}

	for index, comp := range complexes {
		fmt.Printf("%.5f starts at byte position %d\n", comp, index)
	}
}

// \a   U+0007 alert or bell
// \b   U+0008 backspace
// \f   U+000C form feed
// \n   U+000A line feed or newline
// \r   U+000D carriage return
// \t   U+0009 horizontal tab
// \v   U+000b vertical tab
// \\   U+005c backslash
// \'   U+0027 single quote  (valid escape only within rune literals)
// \"   U+0022 double quote  (valid escape only within string literals)
func DisplayRuneLiterals() {
	fmt.Println([]byte("café"))
	fmt.Println([]rune("café"))
	runes := []rune{
		'a',
		'本',
		'\t',
		'\000',
		'\007',
		'\377',
		'\x07',
		'\xff',
		'\u12e4',
		'\U00101234',
		'\'', // rune literals containing single quote character
	}

	for index, word := range runes {
		fmt.Printf("%#U starts at byte position %d\n", word, index)
	}
}

func DisplayStringLiterals() {
	strings := []string{
		`abc`,
		`\n
		\n`,
		"\n",
		"\"",
		"Hello, world!\n",
		"日本語",
		"\u65e5本\U00008a9e",
		`日本語`,
		"\u65e5\u672c\u8a9e",
		"\U000065e5\U0000672c\U00008a9e",
		"\xe6\x97\xa5\xe6\x9c\xac\xe8\xaa\x9e",
	}

	var value string
	for _, value = range strings {
		fmt.Printf("%s", value)
	}
}
