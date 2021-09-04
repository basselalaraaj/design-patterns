package command

import (
	"fmt"
)

type command struct {
	commandType string
	args        []string
}

func CreateCommand(commandType string, args []string) (*command, error) {
	c := command{commandType: commandType, args: args}

	if err := c.decode([]byte{}); err != nil {
		return nil, err
	}

	if err := c.validateCommand(); err != nil {
		return nil, err
	}

	return &c, nil
}

func (c command) decode(data []byte) error {
	// decodes and initialise
	fmt.Println("command is decoded")
	return nil
}

func (c command) validateCommand() error {
	// validates command type
	fmt.Println("command is validated")
	return nil
}
