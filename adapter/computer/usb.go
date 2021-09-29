package computer

import "fmt"

type USB struct{}

func (u *USB) insertIntoUSBPort() {
	fmt.Println("inserting USB cable into USB port!")
}
