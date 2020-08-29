package interfaces

type StockPosition struct {
	ticker     string
	sharePrice float64
	count      float64
}

// method to determine the value of a stock position
func (sp *StockPosition) getValue() float64 {
	return sp.count * sp.sharePrice
}

type Car struct {
	make  string
	model string
	price float64
}

// method to determine the value of a car
func (c *Car) getValue() float64 {
	return c.price
}

type valuable interface {
	getValue() float64
}
