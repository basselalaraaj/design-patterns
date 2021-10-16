package notifier

import "fmt"

type slackNotifier struct {
	notifier notifier
}

func NewSlackNotifier(n notifier) *slackNotifier {
	return &slackNotifier{notifier: n}
}

func (s *slackNotifier) Send(m string) {
	s.notifier.Send(m)
	fmt.Println("send slack message!", m)
}
