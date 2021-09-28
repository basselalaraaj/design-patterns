package main

import (
	"fmt"

	"github.com/basselalaraaj/design-patterns/singleton/product"
)

func main() {
	for i := 0; i < 30; i++ {
		go func() {
			product := product.GetProduct()
			fmt.Println(product.GetName())
		}()
	}

	fmt.Scanln()
}
