package interfaces

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValuable(t *testing.T) {
	var o valuable = &StockPosition{ticker: "ticker", sharePrice: 250.0, count: 10}
	assert.Equal(t, o.getValue(), 2500.0)
	o = &Car{price: 1800}
	assert.Equal(t, o.getValue(), 1800.0)
}
