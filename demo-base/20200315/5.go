package main

import "fmt"

type Benz struct {

}

func (this *Benz) Run() {
	fmt.Println("Banz is running...")
}

type BMW struct {

}

func (this *BMW) Run() {
	fmt.Println("BWN is running...")
}

type Zhang3 struct {

}

func (zhang3 Zhang3) DriveBanZ(benz *Benz) {
	fmt.Println("zhang3 Drive Benz")
	benz.Run()
}

func (zhang3 *Zhang3) DriveBMW(bmw *BMW) {
	fmt.Println("zhang3 drive BMW")
}

type Li4 struct {

}

func (li4 *Li4) DriveBenZ(benz *Benz) {
	fmt.Println("li4 Drive Benz")
	benz.Run()
}

func (li4 *Li4) DriveBMW(bmw *BMW) {
	fmt.Println("li4 Drive BMW")
	bmw.Run()
}

func main() {
	benz := &Benz{}
	zhang3 := &Zhang3{}
	zhang3.DriveBanZ(benz)

	bmw := &BMW{}
	li4 := &Li4{}
	li4.DriveBMW(bmw)
}