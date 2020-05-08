package main

import "fmt"

/*
	A.process -> B.convert()
*/

//IProcess interface
type IProcess interface {
	process(string)
}

//Adapter struct
type Adapter struct {
	AdapterType string
	adaptee Adaptee
}

//Adapter class method process
func (adapter Adapter) process(s string) {
	fmt.Println("Adapter process： ", adapter.AdapterType)
	adapter.adaptee.convert(s)
}

//Adaptee struct
type Adaptee struct {
	adapteeType string
}

//Adaptee class method convert
func (adaptee Adaptee) convert(s string) {
	fmt.Println("Adaptee convert method: ", adaptee.adapteeType)
	if s == "hello" {
		fmt.Println("convert: ", s, "->", "你好")
	}

}

//main method
func main() {
	var processer IProcess = Adapter{
		AdapterType: "hello",
		adaptee:Adaptee{
			adapteeType: "你好",
		},
	}
	processer.process("hello")
}

