package email

import "fmt"

type email struct {
	to, from, subject, body string
}

type EmailBuilder struct {
	email
}

func (b *EmailBuilder) To(to string) *EmailBuilder {
	b.to = to
	return b
}

func (b *EmailBuilder) From(from string) *EmailBuilder {
	b.from = from
	return b
}

func (b *EmailBuilder) Subject(subject string) *EmailBuilder {
	b.subject = subject
	return b
}

func (b *EmailBuilder) Body(body string) *EmailBuilder {
	b.body = body
	return b
}

func SendEmail(cb func(b *EmailBuilder)) {
	emailBuilder := EmailBuilder{}
	cb(&emailBuilder)
	sendEmail(&emailBuilder.email)
}

func sendEmail(email *email) {
	fmt.Printf("To: %v\nFrom: %v\nSubject: %v\n\n%v", email.to, email.from, email.subject, email.body)
}
