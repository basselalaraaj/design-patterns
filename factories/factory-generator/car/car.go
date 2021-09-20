package car

type car struct {
	brand, model, color string
}

func (c *car) Color() string {
	return c.color
}

func (c *car) Brand() string {
	return c.brand
}

func NewCar(brand, model string) *carFactory {
	return &carFactory{
		brand: brand,
		model: model,
	}
}

type carFactory struct {
	brand, model string
}

func (f *carFactory) Create(color string) *car {
	return &car{
		brand: f.brand,
		model: f.model,
		color: color,
	}
}
