package hospital

import "fmt"

type doctor struct {
	nextHandler
}

func (d *doctor) handle(p *patient) {
	fmt.Println("Doctor checking patient")
	d.handleNext(p)
}
