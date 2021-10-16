package order

import "fmt"

type order interface {
	PlaceOrder(items []string) error
}

type orderFacade struct {
	delivery *delivery
	payment  *payment
}

func NewOrderFacade() order {
	return &orderFacade{
		delivery: &delivery{},
		payment:  &payment{},
	}
}

func (o *orderFacade) PlaceOrder(items []string) error {
	fmt.Println("Placing order..")
	fmt.Println(items)
	o.payment.capturePayment()
	o.delivery.deliverPackage()
	fmt.Println("order complete!")
	return nil
}
