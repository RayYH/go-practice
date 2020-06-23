package basic

// 常量声明格式 const identifier [type] = value - Go 支持布尔型、数值型、字符串型常量
// type 在大多数时候都可以省略，因为编译器可以根据变量的值来推断其类型
const Pi = 3.14159

// 因式分解字 - 使用小括号将多个常量放在同一组里
const (
	Sunday    = 0
	Monday    = 1
	Tuesday   = 2
	Wednesday = 3
	Thursday  = 4
	Friday    = 5
	Saturday  = 6
)

// 并行赋值
const name, age = "Ray", 24

// iota 只能用于常量 (const 关键字) 的声明而不能用于其他地方
// iota - 从 0 开始，每次换到 **新的一行** 其值都会自动 + 1
// 这同时表明位于其前面的一行都必须赋值
const (
	a = iota // 0
	b = iota // 1
	c = iota // 2
)

// 可以省略后面的 iota
const (
	d = iota // 0
	e        // 1
	f        // 2
)

//  显式赋值并不会中断 iota 自增
const (
	g = iota     // g = iota = 0
	h            // h = iota = 1
	i = "string" // i = "string", iota = 2
	j            // iota 中断之后不会继续赋予 iota 的值，而是被赋予上一个自定义的的值 j=0="string"，此时 iota = 3
	k = iota     // k = iota = 4
)

const (
	l = 7    // 第一次 iota 前面的常量必须赋值，但此时 iota 已经被初始化为 0
	m = 8    // 此时 iota = 1
	n        // n = m = 8 但是 iota = 2
	o = iota // 3
	p        // 4
)

// 这里 type 关键字的作用是类型别名
type Color int

// 可以使用 _ 跳过一些值
const (
	RED    Color = iota // 0
	ORANGE              // 1
	YELLOW              // 2
	_                   // 跳过值
	_                   // ..
	INDIGO              // 5
	VIOLET              // 6
)

type ByteSize float64

const (
	_           = iota             // 忽视掉 0
	KB ByteSize = 1 << (10 * iota) // 1 << (10*1)
	MB                             // 1 << (10*2)
	GB                             // 1 << (10*3)
	TB                             // 1 << (10*4)
	PB                             // 1 << (10*5)
	EB                             // 1 << (10*6)
	ZB                             // 1 << (10*7)
	YB                             // 1 << (10*8)
)
