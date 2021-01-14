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
	// anonymous fields
	intStruct
	floatStruct
	bs          boolStruct
	floatValue  float64 // override inner struct
	stringValue string
}
