package structs

import (
	"fmt"
	"math"
	"reflect"
	"strings"
)

type SimpleStruct struct {
	a, b int
}

type Interval struct {
	start int
	end   int
}

type Person struct {
	firstName string
	lastName  string
}

type File struct {
	fd   int
	name string
}

type TagType struct {
	Field1 bool   `description:"An important answer"`
	Field2 string `description:"The name of the thing"`
	Field3 int    `description:"How much there are"`
}

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

func RefTag(tt TagType, ix int) {
	ttType := reflect.TypeOf(tt)
	ixField := ttType.Field(ix)
	fmt.Printf("%v\n", ixField.Tag.Get("description"))
}

func upPerson(p *Person) {
	p.firstName = strings.ToUpper(p.firstName)
	p.lastName = strings.ToUpper(p.lastName)
}

func NewFile(fd int, name string) *File {
	if fd < 0 {
		return nil
	}

	return &File{fd, name}
}

// Go 方法是作用在接收者上的一个函数，接收者是某种类型的变量
// 接收者类型可以是几乎任何类型，不仅仅是结构体类型，任何类型都可以有方法，甚至可以是函数类型
// 接收者不能是一个接口类型，因为接口是一个抽象定义，但是方法却是具体实现
// 接收者不能是一个指针类型，但是它可以是任何其他允许类型的指针
// func (recv receiver_type) methodName(parameter_list) (return_value_list) { ... }
type TwoIntegers struct {
	a int
	b int
}

// 接受者 - 方法名 - 返回值
func (ti *TwoIntegers) AddThem() int {
	return ti.a + ti.b
}

func (ti *TwoIntegers) AddToParam(param int) int {
	return ti.AddThem() + param
}

// 非结构体上的方法
type IntVector []int

func (v IntVector) Sum() (s int) {
	for _, value := range v {
		s += value
	}
	return
}

// 函数和方法的区别
// Function(recv) vs recv.Method()
// 指针方法和值方法都可以在指针或非指针上被调用
type IntContainer struct {
	thing int
}

func (b *IntContainer) change()      { b.thing = 1 }
func (b IntContainer) write() string { return fmt.Sprint(b) }

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
