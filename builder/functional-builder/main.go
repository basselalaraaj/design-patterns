package main

import (
	"fmt"

	"github.com/basselalaraaj/design-patterns/builder/functional-builder/person"
)

func main() {
	b := person.NewPersonBuilder()
	p := b.
		Name("john", "doe").
		Age(23).
		Build()

	fmt.Println(p.FullName())
	fmt.Println(p.Age())
}
