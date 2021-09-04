package main

import (
	"fmt"

	"github.com/basselalaraaj/design-patterns/solid-design/open-closed-principle/command"
)

func main() {
	var fooC *command.FooCommand
	err := command.ExecuteCommand(fooC)
	if err != nil {
		fmt.Println(err)
	}
}
