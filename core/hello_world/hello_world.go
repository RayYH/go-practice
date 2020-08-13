package main

import "fmt"

func main() {
	fmt.Println("Hello World!")
}

// 为什么我们需要 Go 语言，Go 语言有哪些特性
// 1. Go 在语言级别支持协程 (goroutine)，使用通道 (channel) 实现了 CSP (Communicating Sequential Process) 模型
//    并发编程模型有两个流派，共享内存和消息传递 (经典代表是 Erlang)
//    关于 Erlang 风格的并发模型，其主体思想有两点：
//    1) 轻量级的进程
//    2) 消息 (通常基于消息队列) 乃进程间通信的唯一方式
// 2. Go 支持自动垃圾回收 (GC)
// 3. Go 语言强制统一代码风格
// 4. Go 语言支持 defer/panic/recover 关键字，基于此 Go 语言在错误处理上有自己独特的实现方式
// 5. Go 语言支持多个返回值
// 6. Go 语言不支持函数和操作符的重载 (不同于 C++, Java, C#)
// 7. Go 语言不支持显式的继承，Go 的继承需要通过组合的方式来实现
// 8. Go 放弃了大量的 OOP 特性，但 Go 语言支持接口
// 9. Go 语言支持匿名函数与闭包
