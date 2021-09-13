package main

import "fmt"

type person struct {
	firstName, lastName string
	age                 int

	// address
	street, city string
	houseNumber  int
}

func newPersonBuilder() *personBuilder {
	return &personBuilder{&person{}}
}

type personBuilder struct {
	person *person
}

func (b *personBuilder) setName(firstName, lastName string) *personBuilder {
	b.person.firstName = firstName
	b.person.lastName = lastName
	return b
}

func (b *personBuilder) setAge(age int) *personBuilder {
	b.person.age = age
	return b
}

func (b *personBuilder) address() *personAddressBuilder {
	return &personAddressBuilder{*b}
}

func (b *personBuilder) build() *person {
	return b.person
}

type personAddressBuilder struct {
	personBuilder
}

func (b *personAddressBuilder) setStreet(street string) *personAddressBuilder {
	b.person.street = street
	return b
}

func (b *personAddressBuilder) setCity(city string) *personAddressBuilder {
	b.person.city = city
	return b
}

func (b *personAddressBuilder) setHouseNumber(houseNumber int) *personAddressBuilder {
	b.person.houseNumber = houseNumber
	return b
}

func main() {
	pb := newPersonBuilder()
	pb.
		setName("john", "doe").
		setAge(29)

	pb.
		address().
		setStreet("marsstreet").
		setHouseNumber(14).
		setCity("rotterdam")

	person := pb.build()
	fmt.Println("first name", person.firstName)
	fmt.Println("age", person.age)
	fmt.Println("street", person.street)
	fmt.Println("city", person.city)
}
