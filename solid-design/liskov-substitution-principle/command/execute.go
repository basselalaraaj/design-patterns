package command

type command interface {
	execute()
}

func Execute(c command) {
	c.execute()
}
