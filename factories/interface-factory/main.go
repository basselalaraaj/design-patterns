package main

import (
	"fmt"

	"github.com/basselalaraaj/design-patterns/factories/interface-factory/person"
)

func main() {
	p := person.NewPerson("john", "doe", 20)

	fmt.Println(p.Name())
}
