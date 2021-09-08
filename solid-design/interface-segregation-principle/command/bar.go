package command

import "fmt"

type Bar struct {
}

func (b Bar) execute() {
	fmt.Println("bar run!")
}
