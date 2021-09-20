package person

import "fmt"

func NewPerson(firstName, lastName string, age int) *person {
	return &person{
		firstName: firstName,
		lastName:  lastName,
		age:       age,
		active:    1,
	}
}

type person struct {
	firstName, lastName string
	age, active         int
}

func (p *person) Name() string {
	return fmt.Sprintf("%s %s", p.firstName, p.lastName)
}
