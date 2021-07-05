package basic_types

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// 1. Go 语言支持整型、浮点型、复数类型
// 2. int 和 uint 类型在 32/64 位操作系统上的长度为 32 位/64 位，uintptr 的长度被设定为足够存放一个指针即可
// 3. Go 语言中没有 float 类型和 double 类型，取而代之的是 float32 和 float64
// 4. int 型是计算最快的一种类型
// 5. 整型的零值为 0，浮点型的零值为 0.0

// 类型清单
// uint8       the set of all unsigned  8-bit integers (0 to 255)
// uint16      the set of all unsigned 16-bit integers (0 to 65535)
// uint32      the set of all unsigned 32-bit integers (0 to 4294967295)
// uint64      the set of all unsigned 64-bit integers (0 to 18446744073709551615)
//
// int8        the set of all signed  8-bit integers (-128 to 127)
// int16       the set of all signed 16-bit integers (-32768 to 32767)
// int32       the set of all signed 32-bit integers (-2147483648 to 2147483647)
// int64       the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)
//
// float32     the set of all IEEE-754 32-bit floating-point numbers
// float64     the set of all IEEE-754 64-bit floating-point numbers
//
// complex64   the set of all complex numbers with float32 real and imaginary parts
// complex128  the set of all complex numbers with float64 real and imaginary parts
//
// byte        alias for uint8
// rune        alias for int32, represents a Unicode code point
//
// uint        either 32 or 64 bits
// int         same size as uint
// uintptr     an unsigned integer large enough to store the uninterpreted bits of a pointer value

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
