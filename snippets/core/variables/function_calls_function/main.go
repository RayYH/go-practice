package main

var a string

func f2() {
	// 始终打印全局变量 a
	print(a)
}

func f1() {
	// 声明了局部变量 a 并进行打印
	a := "O"
	print(a)

	// 这里虽然调用的是 f2 但是 f2 并不是一个闭包，而是一个普通函数，因此 f2 打印的仍是全局变量的值
	f2()

	// 这里 f3 是一个闭包，也就是说与外围的作用域相同，因此访问的 a 是 f1 函数的局部变量 a
	f3 := func() {
		print(a)
	}
	f3()
}

func main() {
	a = "G"
	print(a)
	f1()
}
