package order

import "fmt"

type delivery struct {
}

func (d *delivery) deliverPackage() {
	fmt.Println("delivering package to the customer!")
}
