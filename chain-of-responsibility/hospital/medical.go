package hospital

import "fmt"

type medical struct {
	nextHandler
}

func (m *medical) handle(p *patient) {
	fmt.Println("Medical giving medicine to patient")
	m.handleNext(p)
}
