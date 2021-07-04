package main

// #include <stdio.h>
// #include <stdlib.h>
import "C"
import "unsafe"

func Print(s string) {
	// C 当中并没有明确的字符串类型，如果你想要将一个 string 类型的变量从 Go 转换到 C 时，可以使用 C.CString(s)
	cs := C.CString(s)
	// Go 的内存管理机制无法管理通过 C 代码分配的内存，开发人员需要通过手动调用 C.free 来释放变量的内存
	defer C.free(unsafe.Pointer(cs))
	C.fputs(cs, (*C.FILE)(C.stdout))
}

func main() {
	Print("Hello World")
}
