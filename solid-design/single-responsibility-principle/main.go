package main

import (
	"fmt"

	"github.com/basselalaraaj/design-patterns/solid-design/single-responsibility-principle/command"
)

func main() {
	c, err := command.CreateCommand("foo", []string{"test"})
	if err != nil {
		fmt.Println(err)
	}

	err = command.ExecuteCommand(c)
	if err != nil {
		fmt.Println(err)
	}
}
