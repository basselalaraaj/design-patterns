package product

import (
	"fmt"
	"sync"
)

type product struct{}

var once sync.Once
var instance *product

func GetProduct() *product {
	once.Do(func() {
		instance = &product{}
		fmt.Println("created new instance")
	})

	return instance
}

func (p *product) GetName() string {
	return "abc"
}
