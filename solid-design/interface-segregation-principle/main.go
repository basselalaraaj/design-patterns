package main

import "github.com/basselalaraaj/design-patterns/solid-design/interface-segregation-principle/command"

func main() {
	bar := command.Bar{}
	command.Execute(bar)

	foo := command.Foo{}
	command.Execute(foo)
	command.Validate(foo)
	command.ValidatedAndExecute(foo)
}
