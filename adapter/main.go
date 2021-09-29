package main

import (
	"github.com/basselalaraaj/design-patterns/adapter/computer"
)

func main() {
	mac := computer.Mac{}

	usbc := computer.USBC{}
	mac.InsertCable(&usbc)

	usb := computer.USB{}
	usbAdapter := computer.Adapter{
		USB: &usb,
	}
	mac.InsertCable(&usbAdapter)
}
