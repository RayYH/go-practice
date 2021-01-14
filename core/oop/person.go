package oop

import "strings"

type Person struct {
	firstName string
	lastName  string
}

func (p *Person) FirstName() string {
	return p.firstName
}

func (p *Person) SetFirstName(aFirstName string) {
	p.firstName = aFirstName
}

func (p *Person) LastName() string {
	return p.lastName
}

func (p *Person) SetLastName(aLastName string) {
	p.lastName = aLastName
}

func upperPerson(p *Person) {
	p.firstName = strings.ToUpper(p.firstName)
	p.lastName = strings.ToUpper(p.lastName)
}
