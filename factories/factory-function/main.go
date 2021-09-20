package main

import (
	"fmt"

	"github.com/basselalaraaj/design-patterns/factories/factory-function/person"
)

func main() {
	p := person.NewPerson("john", "doe", 20)

	fmt.Println(p.Name())
}
