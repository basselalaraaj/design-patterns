package person

import "fmt"

type person struct {
	firstName, lastName string
	age                 int
}

func (p *person) FullName() string {
	return fmt.Sprintf("%s %s", p.firstName, p.lastName)
}

func (p *person) Age() int {
	return p.age
}

func NewPersonBuilder() *personBuilder {
	return &personBuilder{}
}

type personMod func(p *person)
type personBuilder struct {
	handlers []personMod
}

func (b *personBuilder) Name(firstName, lastName string) *personBuilder {
	b.handlers = append(b.handlers, func(p *person) {
		p.firstName = firstName
		p.lastName = lastName
	})
	return b
}

func (b *personBuilder) Age(age int) *personBuilder {
	b.handlers = append(b.handlers, func(p *person) {
		p.age = age
	})
	return b
}

func (b *personBuilder) Build() *person {
	p := person{}
	for _, h := range b.handlers {
		h(&p)
	}
	return &p
}
