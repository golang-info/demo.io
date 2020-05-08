package main

import (
	"fmt"
)

type DataWriter interface {
	WriteData(data interface{}) error
	PrintData(data interface{}) error
}

type file struct {
}

func (d *file) WriteData(data interface{}) error {
	fmt.Println("WriteData: ", data)
	return nil
}

func (d *file) PrintData(data interface{}) error {
	fmt.Println("PrintData: ", data)
	return nil
}

func main() {
	f := new(file)

	var writer DataWriter

	writer = f

	writer.WriteData("dataWrite")
	writer.PrintData("dataPrint")
}
