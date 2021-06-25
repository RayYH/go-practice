package main

var a = "G"

func n() {
	print(a)
}

func m() {
	// 这里直接使用的 = 与 006 中的 := 不同，这里修改了全局变量的值
	a = "O"
	print(a)
}

func main() {
	n() // "G"
	m() // "O"
	n() // "O"
}
