package command

import "fmt"

func ExecuteCommand(c *command) error {
	switch c.commandType {
	case "foo":
		return FooCommand()
	default:
		fmt.Println("command doesn't exist")
	}
	return nil
}
