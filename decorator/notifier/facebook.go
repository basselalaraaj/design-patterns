package notifier

import "fmt"

type facebookkNotifier struct {
	notifier notifier
}

func NewFacebookkNotifier(n notifier) *facebookkNotifier {
	return &facebookkNotifier{notifier: n}
}

func (f *facebookkNotifier) Send(m string) {
	f.notifier.Send(m)
	fmt.Println("send facebook message!", m)
}
