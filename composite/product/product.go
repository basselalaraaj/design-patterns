package product

type price float32

type product struct {
	price price
}

func NewProduct(price price) *product {
	return &product{price: price}
}

func (p *product) GetPrice() price {
	return p.price
}
