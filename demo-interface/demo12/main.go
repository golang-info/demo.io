package main

import "fmt"

type Phone interface {
	call()
}

type NokiaPhone struct {
}

func (n NokiaPhone) call() {
	fmt.Println("I am Nokia, I can call you.")
}

type ApplePhone struct {
}

func (i ApplePhone) call() {
	fmt.Println("I am Phone, I can call you.")
}

func main() {
	var phone Phone
	phone = new(NokiaPhone)
	phone.call()

	phone = new(ApplePhone)
	phone.call()
}
