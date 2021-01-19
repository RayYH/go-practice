package oop

import (
	"fmt"
	"strings"
)

type Person struct {
	firstName string
	lastName  string
}

// Getters
func (p *Person) FirstName() string {
	return p.firstName
}

func (p *Person) LastName() string {
	return p.lastName
}

// Setters
func (p *Person) SetFirstName(aFirstName string) {
	p.firstName = aFirstName
}

func (p *Person) SetLastName(aLastName string) {
	p.lastName = aLastName
}

// Constructor
func NewPerson(firstName, lastName string) *Person {
	return &Person{firstName: firstName, lastName: lastName}
}

func upperPerson(p *Person) {
	p.firstName = strings.ToUpper(p.firstName)
	p.lastName = strings.ToUpper(p.lastName)
}

func (p Person) String() string {
	return fmt.Sprintf("%s %s", p.firstName, p.lastName)
}
