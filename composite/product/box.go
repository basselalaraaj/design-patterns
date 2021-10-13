package product

type box struct {
	components []component
}

func NewBox() *box {
	return &box{}
}

func (b *box) Add(c component) {
	b.components = append(b.components, c)
}

func (b *box) GetPrice() price {
	var p price
	for _, c := range b.components {
		p += c.GetPrice()
	}
	return p
}
