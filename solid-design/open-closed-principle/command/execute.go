package command

import (
	"errors"
	"fmt"
)

type Command interface {
	Execute() error
}

func ExecuteCommand(c Command) error {
	fmt.Println("start executing the command")
	err := c.Execute()
	if err != nil {
		return errors.New("something went wrong")
	}
	return nil
}
