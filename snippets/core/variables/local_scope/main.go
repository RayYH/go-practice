package main

var a = "G"

func printA() {
	// 由于该函数体内没有声明变量 a，因此这里的 a 指的就是全局变量 a
	print(a)
}

func modifyAndPrintA() {
	// 这里相当于重新定义一个变量 a，该局部变量会在函数退出后销毁掉
	a := "O"
	print(a)
}

func main() {
	printA()          // "G"
	modifyAndPrintA() // "O"
	printA()          // "G"
}
