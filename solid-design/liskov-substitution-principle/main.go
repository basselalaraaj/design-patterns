package main

import "github.com/basselalaraaj/design-patterns/solid-design/liskov-substitution-principle/command"

func main() {
	foo := command.Foo{}
	command.Execute(foo)

	bar := command.Bar{}
	command.Execute(bar)
}
