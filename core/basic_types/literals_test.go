package basic_types

// https://golang.org/ref/spec#Integer_literals

func ExampleDisplayIntegerLiterals() {
	DisplayIntegerLiterals()
	// Output: 42 42 384 384 384 195951310 195951310 113774485586118 1701411834604692310 17014118346046930
}

func ExampleDisplayFloatPointLiterals() {
	DisplayFloatPointLiterals()
	// Output: 0.000 72.400 72.400 2.718 1.000 0.000 1000000.000 0.250 12345.000 15.000 15.000 0.250 2048.000 1.938 0.500 0.125 348.000
}

func ExampleDisplayImaginaryLiterals() {
	DisplayImaginaryLiterals()
	// Output: (0.00000+0.00000i) starts at byte position 0
	// (0.00000+123.00000i) starts at byte position 1
	// (0.00000+83.00000i) starts at byte position 2
	// (0.00000+2748.00000i) starts at byte position 3
	// (0.00000+0.00000i) starts at byte position 4
	// (0.00000+2.71828i) starts at byte position 5
	// (0.00000+1.00000i) starts at byte position 6
	// (0.00000+0.00000i) starts at byte position 7
	// (0.00000+1000000.00000i) starts at byte position 8
	// (0.00000+0.25000i) starts at byte position 9
	// (0.00000+12345.00000i) starts at byte position 10
	// (0.00000+0.25000i) starts at byte position 11
	// (100.00000+10.00000i) starts at byte position 12
}

func ExampleDisplayRuneLiterals() {
	DisplayRuneLiterals()
	// Output: [99 97 102 195 169]
	// [99 97 102 233]
	// U+0061 'a' starts at byte position 0
	// U+672C '本' starts at byte position 1
	// U+0009 starts at byte position 2
	// U+0000 starts at byte position 3
	// U+0007 starts at byte position 4
	// U+00FF 'ÿ' starts at byte position 5
	// U+0007 starts at byte position 6
	// U+00FF 'ÿ' starts at byte position 7
	// U+12E4 'ዤ' starts at byte position 8
	// U+101234 starts at byte position 9
	// U+0027 ''' starts at byte position 10
}

func ExampleDisplayStringLiterals() {
	DisplayStringLiterals()
	// Output: abc\n
	//		\n
	//"Hello, world!
	//日本語日本語日本語日本語日本語日本語
}
