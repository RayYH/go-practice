package oop

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCustomerString(t *testing.T) {
	c := new(Customer)
	c.Name = "Ray Hong"
	c.Log = Log{"first message"}
	// customer can call Add method directly
	// since Log is anonymous
	c.Add("second message")
	assert.Equal(t, "Ray Hong\nLog:{first message\nsecond message}\n", fmt.Sprint(c))
}

func TestCustomerHoldsRefLog(t *testing.T) {
	c1 := new(CustomerHoldsRef)
	c1.Name = "Ray Hong"
	c1.log = new(Log)
	c1.log.msg = "first message"
	assert.Equal(t, fmt.Sprint(c1.Log()), "first message")
	c2 := CustomerHoldsRef{
		Name: "Ray Hong",
		log:  &Log{"first message"},
	}
	c2.Log().Add("second message")
	assert.Equal(t, fmt.Sprint(c2.Log()), "first message\nsecond message")
}
