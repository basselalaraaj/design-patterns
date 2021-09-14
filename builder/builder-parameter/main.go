package main

import (
	"github.com/basselalaraaj/design-patterns/builder/builder-parameter/email"
)

func main() {
	email.SendEmail(func(b *email.EmailBuilder) {
		b.
			To("info@example.com").
			From("test@example.com").
			Subject("Welcome!").
			Body("Welkom home everybody!")
	})
}
