package computer

import "fmt"

type cable interface {
	insertIntoUSBCPort()
}

type Mac struct{}

func (c *Mac) InsertCable(cable cable) {
	fmt.Println("inserting cable")
	cable.insertIntoUSBCPort()
}
