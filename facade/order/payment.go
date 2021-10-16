package order

import "fmt"

type payment struct {
}

func (p *payment) capturePayment() {
	fmt.Println("capture payment!")
}
