package main

import (
	"fmt"
	"reflect"
)

type UnKnownType struct {
	s1, s2, s3 string
}

func (n UnKnownType) String() string {
	return n.s1 + " - " + n.s2 + " - " + n.s3
}

var secret interface{} = UnKnownType{"Ada", "Go", "Oberon"}

func main() {
	value := reflect.ValueOf(secret)

	// 除了使用 reflect.TypeOf 方法获取指定变量的类型外，我们还可以使用 value.Type()
	typ := reflect.TypeOf(secret)
	fmt.Println(typ)

	// Kind() 返回变量的类型
	// type Kind uint
	//
	// const (
	//	 Invalid Kind = iota
	//	 Bool
	//	 Int
	//	 Int8
	//	 Int16
	//	 Int32
	//	 Int64
	//	 Uint
	//	 Uint8
	//	 Uint16
	//	 Uint32
	//	 Uint64
	//	 Uintptr
	//	 Float32
	//	 Float64
	//	 Complex64
	//	 Complex128
	//	 Array
	//	 Chan
	//	 Func
	//	 Interface
	//	 Map
	//	 Ptr
	//	 Slice
	//	 String
	//	 Struct
	//	 UnsafePointer
	// )
	knd := value.Kind()
	fmt.Println(knd)

	// NumField() 返回结构体中的字段个数
	for i := 0; i < value.NumField(); i++ {
		// Field() 返回结构体指定索引处的字段值
		fmt.Printf("Field %d: %v\n", i, value.Field(i))
		// 如果我们使用 value.Field(i).SetString("Something") 来尝试修改字段值则会触发一个 panic
	}

	// 调用第一个方法，即 String()
	results := value.Method(0).Call(nil)
	fmt.Println(results)
}
