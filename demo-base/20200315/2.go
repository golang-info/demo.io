package main

import "fmt"

type Banker struct {

}

func (this *Banker) Save() {
	fmt.Println("进行了 存款业务。。。")
}

func (this *Banker) Transfer() {
	fmt.Println("进行了 转账业务。。。")
}

func (this *Banker) Pay() {
	fmt.Println("进行了 支付业务。。。")
}

func main() {
	banker := &Banker{}

	banker.Save()
	banker.Transfer()
	banker.Pay()
}