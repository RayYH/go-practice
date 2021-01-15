package oop

type Interval struct {
	start int
	end   int
}

// 1. A method is a function with a special receiver argument.
// 2. You can only declare a method with a receiver whose type is defined in the same package as the method.
// 3. Methods with pointer receivers can modify the value to which the receiver points.
//    Since methods often need to modify their receiver, pointer receivers are more common than value receivers.
// func (recv receiver_type) methodName(parameter_list) (return_value_list) { ... }
func (i *Interval) duration() int {
	return i.end - i.start
}

func (i *Interval) durationInMillSeconds() int {
	return i.duration() * 1000
}
