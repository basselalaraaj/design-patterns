package main

import (
	"fmt"

	"github.com/basselalaraaj/design-patterns/prototype/house"
)

func main() {
	houseA := house.NewHouse("utenstraat", "holland", 43)
	fmt.Println("houseA", houseA)

	houseB := houseA.Clone()
	houseB.Pool = true
	houseB.Garden = true

	fmt.Println("houseB", houseB)

	fmt.Println("houseA", houseA)
}
