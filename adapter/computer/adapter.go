package computer

import "fmt"

type Adapter struct {
	USB *USB
}

func (a *Adapter) insertIntoUSBCPort() {
	fmt.Println("adding USB to USBC adapter")
	a.USB.insertIntoUSBPort()
}
