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

func main() {
	sb := &SaveBanker{}
	sb.DoBusi()

	tb := &TransferBanker{}
	tb.DoBusi()

	pb := &PayBanker{}
	pb.DoBusi()
}


