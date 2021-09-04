package command

import "fmt"

func FooCommand() error {
	fmt.Println("foo it is!")
	return nil
}
