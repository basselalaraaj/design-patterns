package main

import (
	"fmt"

	"github.com/basselalaraaj/design-patterns/facade/order"
)

func main() {
	orderFacade := order.NewOrderFacade()

	err := orderFacade.PlaceOrder([]string{
		"phone",
		"watch",
		"wallet",
	})
	if err != nil {
		fmt.Println(err)
	}
}
