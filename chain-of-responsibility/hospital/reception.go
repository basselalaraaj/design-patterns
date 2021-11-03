package hospital

import "fmt"

type reception struct {
	nextHandler
}

func (r *reception) handle(p *patient) {
	fmt.Println("Reception registering patient")
	r.handleNext(p)
}
