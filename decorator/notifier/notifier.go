package notifier

import "fmt"

type notifier interface {
	Send(m string)
}

type defaultNotifier struct {
}

func NewNotifier() *defaultNotifier {
	return &defaultNotifier{}
}

func (d *defaultNotifier) Send(m string) {
	fmt.Println("send default message!", m)
}
