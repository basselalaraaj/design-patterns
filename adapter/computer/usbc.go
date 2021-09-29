package computer

import "fmt"

type USBC struct{}

func (u *USBC) insertIntoUSBCPort() {
	fmt.Println("inserting USBC cable into USBC port!")
}
