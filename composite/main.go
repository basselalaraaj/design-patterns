package main

import (
	"fmt"

	"github.com/basselalaraaj/design-patterns/composite/product"
)

func main() {
	// create level 1 box
	box1 := product.NewBox()
	product1 := product.NewProduct(5)
	box1.Add(product1)

	// create level 2 box
	box2 := product.NewBox()
	product2 := product.NewProduct(3)
	product3 := product.NewProduct(1.5)
	box2.Add(product2)
	box2.Add(product3)

	// create level 3 box
	box3 := product.NewBox()
	product4 := product.NewProduct(8.4)
	box3.Add(product4)

	// add level 3 box to level 1 box
	box2.Add(box3)

	// add level 2 box to level 1 box
	box1.Add(box2)

	// show total of box level 1
	fmt.Println("total price box 1:", box1.GetPrice())
}
