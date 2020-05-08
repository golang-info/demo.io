package main

import "fmt"

type Phone interface {
	Call()
}

type NokiaPhone struct {

}

func (nokiaPhone NokiaPhone) Call() {
	fmt.Println("I am Nokia, I can call you!")
}

type ApplePhone struct {

}

func (applePhone ApplePhone) Call() {
	fmt.Println("I am Apply Phone, I can call you!")
}

func main() {
	var phone Phone
	phone = new(NokiaPhone)
	phone.Call()

	phone = new(ApplePhone)
	phone.Call()
}