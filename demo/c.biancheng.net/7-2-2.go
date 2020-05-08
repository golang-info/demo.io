package main

import (
	"fmt"
)

type DataWriter interface {
	WriteData(data interface{}) error
	PrintData(data interface{}) error
}

type file struct {
	D string
}

func (d *file) WriteData(data interface{}) error {
	fmt.Println("WriteData: ", data)
	fmt.Println("WriteData: ", d.D)
	return nil
}

func (d *file) PrintData(data interface{}) error {
	fmt.Println("PrintData: ", data)
	fmt.Println("PrintData: ", d.D)
	return nil
}

func main() {
	f := new(file)
	//f.D = "test"
	f = &file{
		D: "test",
	}

	var writer DataWriter

	writer = f

	writer.WriteData("dataWrite")
	writer.PrintData("dataPrint")
}
