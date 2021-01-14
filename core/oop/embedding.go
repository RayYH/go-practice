package oop

type intStruct struct {
	intValue int
}

type boolStruct struct {
	boolValue bool
}

type floatStruct struct {
	floatValue float64
}

type outerStruct struct {
	intStruct              // anonymous struct field intStruct
	floatStruct            // anonymous struct field floatStruct
	bs          boolStruct // named struct field
	floatValue  float64    // override inner struct
	stringValue string     // self field
}
