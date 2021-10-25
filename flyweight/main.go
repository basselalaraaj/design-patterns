package main

import (
	"fmt"

	"github.com/basselalaraaj/design-patterns/flyweight/user"
	"github.com/brianvoe/gofakeit/v6"
)

func main() {
	game := user.NewGame()
	i := 0
	for i < 10000 {
		p := gofakeit.Person()
		u := game.CreateUser(fmt.Sprintf("%v %v", p.FirstName, p.LastName))
		fmt.Println(u.GetFullName())
		i++
	}

	fmt.Printf("created %v users\n", i)
	fmt.Printf("%v names saved in memory\n", game.CountNames())
}
