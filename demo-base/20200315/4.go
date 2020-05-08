package main

import "fmt"

//抽象的银行业务员
type AbstractBanker interface {
	DoBusi()  //抽象的处理业务接口
}

type SaveBanker struct {

}

func (sb *SaveBanker) DoBusi() {
	fmt.Println("进行了存款")
}

type TransferBanker struct {

}

func (tb *TransferBanker) DoBusi() {
	fmt.Println("进行了转账")
}

type PayBanker struct {

}

func (ph *PayBanker) DoBusi() {
	fmt.Println("进行了支付")
}

//实现架构层（基于抽象层进行业务封装-针对interface接口进行封装）
func BankerBusiness(banker AbstractBanker) {
	//通过接口来向下调用，(多态现象)
	banker.DoBusi()
}

func main() {
	BankerBusiness(&SaveBanker{})
	BankerBusiness(&TransferBanker{})
	BankerBusiness(&PayBanker{})
}


