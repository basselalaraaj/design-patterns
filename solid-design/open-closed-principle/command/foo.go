package command

import "fmt"

type FooCommand struct{}

func (f *FooCommand) Execute() error {
	fmt.Println("foo it is!")
	return nil
}
