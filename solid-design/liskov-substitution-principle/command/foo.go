package command

import "fmt"

type Foo struct {
}

func (f Foo) execute() {
	fmt.Println("foo run!")
}
