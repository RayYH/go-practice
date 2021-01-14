package structs

import (
	"fmt"
	"math"
)

type innerS struct {
	in1 int
	in2 int
}

type outerS struct {
	b      int
	c      float64
	int    // anonymous field
	innerS // anonymous field
	// Go 语言通过组合来达到继承的效果
}

type A struct {
	ax, ay int
}

type B struct {
	A
	bx, by float64
}

type innerStruct struct {
	a, b int
}

type outerStruct struct {
	is innerStruct
	// 当两个字段拥有相同的名字，外层结构体中的名字会覆盖内层结构体中的名字
	a, c int
}


// 在 Go 中，代码复用通过组合和委托实现，多态通过接口的使用来实现
// 当一个匿名类型被内嵌在结构体中时，匿名类型的可见方法也同样被内嵌，这在效果上等同于外层类型继承了这些方法。
// 将父类型放在子类型中的机制提供了一种简单的方式来模拟经典面向对象语言中的子类和继承

type Point struct {
	x, y float64
}

func (p *Point) Abs() float64 {
	return math.Sqrt(p.x*p.x + p.y*p.y)
}

type NamedPoint struct {
	Point
	name string
}

type Log struct {
	msg string
}

type CustomerWithRef struct {
	Name string
	log  *Log
}

type Customer struct {
	Name string
	Log
}

func (l *Log) Add(s string) {
	l.msg += "\n" + s
}

func (l *Log) String() string {
	return l.msg
}

func (c *CustomerWithRef) Log() *Log {
	return c.log
}

func (c *Customer) String() string {
	return c.Name + "\nLog:" + fmt.Sprintln(c.Log)
}

type Camera struct{}

func (c *Camera) TakeAPicture() string {
	return "Click"
}

type Phone struct{}

func (p *Phone) Call() string {
	return "Ring Ring"
}

// 多重继承
type CameraPhone struct {
	Camera
	Phone
}
