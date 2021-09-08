package command

type executer interface {
	execute()
}

type validator interface {
	validate()
}

type command interface {
	executer
	validator
}

func Validate(c validator) {
	c.validate()
}

func Execute(c executer) {
	c.execute()
}

func ValidatedAndExecute(c command) {
	c.validate()
	c.execute()
}
