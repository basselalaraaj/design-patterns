package command

import "fmt"

type Foo struct {
}

func (f Foo) validate() {
	fmt.Println("foo validated!")
}

func (f Foo) execute() {
	fmt.Println("foo run!")
}
