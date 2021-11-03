package main

import "github.com/basselalaraaj/design-patterns/chain-of-responsibility/hospital"

func main() {
	p := hospital.NewPatient(
		"John",
		22,
		[]string{
			"Fever",
			"Cough",
		},
	)
	hospital.VisitDoctor(p)
}
