package main

import (
	"fmt"

	"github.com/basselalaraaj/design-patterns/factories/prototype-factory/employee"
)

func main() {
	developer := employee.NewEmployee(employee.Developer)
	developer.SetName("John")
	fmt.Println(developer)

	manager := employee.NewEmployee(employee.Manager)
	manager.SetName("Greg")
	fmt.Println(manager)

	boss := employee.NewEmployee(employee.Boss)
	boss.SetName("Adam")
	fmt.Println(boss)
}
