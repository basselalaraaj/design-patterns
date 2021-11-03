package hospital

import "fmt"

type cashier struct {
	nextHandler
}

func (c *cashier) handle(p *patient) {
	fmt.Println("Cashier getting money from patient")
	c.handleNext(p)
}
