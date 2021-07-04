package main

var a = "G"

func n() {
	print(a)
}

func m() {
	// 这里直接修改了全局变量 a 的值
	a = "O"
	print(a)
}

func main() {
	n() // "G"
	m() // "O"
	n() // "O"
}
