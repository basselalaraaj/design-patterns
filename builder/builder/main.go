package main

import (
	"github.com/basselalaraaj/design-patterns/builder/builder/product"
)

func main() {
	b := product.NewProductBuilder()
	p := b.SetColor("red").SetSize("xs").Build()

	p.Specification()
}
