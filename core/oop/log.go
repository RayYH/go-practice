package oop

import "fmt"

type Log struct{ msg string }

type Customer struct {
	Name string
	Log
}

type CustomerHoldsRef struct {
	Name string
	log  *Log
}

func (l *Log) Add(s string) {
	l.msg += "\n" + s
}

func (l *Log) String() string {
	return l.msg
}

func (c *Customer) String() string {
	return c.Name + "\nLog:" + fmt.Sprintln(c.Log)
}

func (cwr *CustomerHoldsRef) Log() *Log {
	return cwr.log
}
