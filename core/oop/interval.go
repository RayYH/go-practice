package oop

type Interval struct {
	start int
	end   int
}

// func (recv receiver_type) methodName(parameter_list) (return_value_list) { ... }
func (i *Interval) duration() int {
	return i.end - i.start
}

func (i *Interval) durationInMillSeconds() int {
	return i.duration() * 1000
}
