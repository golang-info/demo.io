package main

import (
	"fmt"
)

//go 语言任何类型都已经实现了空接口
type Empty interface {

}

type USB interface {
	Name() string
	Connector
}

type Connector interface {
	Connect()
}

type PhoneConnector struct {
	name string
}

func (pc PhoneConnector) Name() string {
	return pc.name
}

func (pc PhoneConnector) Connect() {
	fmt.Println("Connected:", pc.name)
}

func main() {
	var a USB
	a = PhoneConnector{"PhoneConnector"}
	a.Connect()
	Disconnect(a)
}

/* func Disconnect(usb USB) {
	//简单类型断言
	if pc, ok := usb.(PhoneConnector); ok {
		fmt.Println("Disconnected: ", pc.name)
		return
	}
	fmt.Println("Unknown decive.")
}  */


func Disconnect(usb interface{}) {
	switch v := usb.(type) {
	case PhoneConnector:
		fmt.Println("Disconnected: ", v.name)
	default:
		fmt.Println("Unknown decive.")
	}
} 