package main

import (
	"github.com/basselalaraaj/design-patterns/decorator/notifier"
)

func main() {
	n := notifier.NewNotifier()
	s := notifier.NewSlackNotifier(n)
	f := notifier.NewFacebookkNotifier(s)
	f.Send("alert!")
}
