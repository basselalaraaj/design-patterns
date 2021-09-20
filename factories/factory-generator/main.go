package main

import (
	"fmt"

	"github.com/basselalaraaj/design-patterns/factories/factory-generator/car"
	"github.com/basselalaraaj/design-patterns/factories/factory-generator/employee"
)

func main() {
	// functional approach
	developerFactory := employee.NewEmployee("developer", 70000)

	developerA := developerFactory("John")
	developerB := developerFactory("Greg")

	fmt.Println(developerA.Name())
	fmt.Println(developerA.Position())

	fmt.Println(developerB.Name())
	fmt.Println(developerB.Position())

	// structural approach
	carFactory := car.NewCar("tesla", "model 3")

	carA := carFactory.Create("black")
	carB := carFactory.Create("White")

	fmt.Println("car A", carA.Color(), carA.Brand())
	fmt.Println("car B", carB.Color(), carB.Brand())
}
