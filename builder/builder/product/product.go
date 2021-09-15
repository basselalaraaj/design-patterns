package product

import "fmt"

type product struct {
	size, color string
}

func (p *product) Specification() {
	fmt.Println("size", p.size)
	fmt.Println("color", p.color)
}

func NewProductBuilder() *productBuilder {
	return &productBuilder{product: &product{}}
}

type productBuilder struct {
	product *product
}

func (b *productBuilder) SetSize(size string) *productBuilder {
	b.product.size = size
	return b
}

func (b *productBuilder) SetColor(color string) *productBuilder {
	b.product.color = color
	return b
}

func (b *productBuilder) Build() *product {
	return b.product
}
