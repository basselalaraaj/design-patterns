package house

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

func NewHouse(street, country string, houseNumber int) *house {
	return &house{
		Street:      street,
		Country:     country,
		HouseNumber: houseNumber,

		Rooms:     4,
		Bathrooms: 2,
		Pool:      false,
		Garden:    false,
	}
}

type house struct {
	Rooms, Bathrooms int
	Pool, Garden     bool

	HouseNumber     int
	Street, Country string
}

func (h *house) Clone() *house {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	err := e.Encode(h)
	if err != nil {
		fmt.Println(err)
	}

	d := gob.NewDecoder(&b)
	result := house{}
	err = d.Decode(&result)
	if err != nil {
		fmt.Println(err)
	}

	return &result
}
