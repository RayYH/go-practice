package types

type Integer int

func (a Integer) LessThan(b Integer) bool {
	return a < b
}

func LessThan(a, b Integer) bool {
	return a < b
}

func (a *Integer) Add(b Integer) {
	*a += b
}
