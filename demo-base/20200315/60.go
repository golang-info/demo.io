package main

import "fmt"

//=========> 抽象层 <==============
type Car interface {
	Run()
}

type Driver interface {
	Drive(car Car)
}

//==========> 实现层 <==============
type BenZ struct {

}

func (benz *BenZ) Run() {
	fmt.Println("Benz is runnint...")
}

type BMW struct {

}

func (bmw *BMW) Run() {
	fmt.Println("Bmw is running")
}

type Zhang_3 struct {

}

func (zhang3 *Zhang_3) Drive(car Car) {
	fmt.Println("Zhang3 drive car")
	car.Run()
}

func Li_4 struct {

}

func (li4 *Li_4) Drive(car Car) {
	fmt.Println("Li4 drive car")
	car.Run()
}

//============> 业务逻辑层 <================
func main() {
	var bmw Car
	bmw = &BMW{}

	var zhang3 Driver
	zhang3 = &Zhang_3{}

	zhang3.Drive(bmw)

	var benz Car
	benz = &BenZ{}

	var li4 Driver
	li4 = &Li_4{}

	li4.Drive(benz)
}




