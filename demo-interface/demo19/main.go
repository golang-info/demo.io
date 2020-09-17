package main

/*
https://mp.weixin.qq.com/s?__biz=MzAxMTA4Njc0OQ==&mid=2651441312&idx=3&sn=f9e0bd59c3b338bf37262c8c2c5e30c6&chksm=80bb1652b7cc9f448f2381218a334ca96d49b0014c201b32158abaa8d019dbf264963f0a31e6&scene=21#wechat_redirect
*/

import "fmt"

type AllInOne interface {
	Phone
	gameConsole
}

// 接口定义
type Phone interface {
	Call()
	SendMsg(msg string) bool
}

type gameConsole interface {
	Play()
}

// 接口实现
type AndroidPhone struct {
	Name string
}

func (a AndroidPhone) Call() {
	fmt.Printf("%s calling...\n", a.Name)
}

func (a AndroidPhone) SendMsg(msg string) bool {
	fmt.Printf("%s sending msg: %s.\n", a.Name, msg)
	return true
}

func (a AndroidPhone) Play() {
	fmt.Println("Play game.")
}

func main() {
	// 接口声明
	//var phone Phone
	// fmt.Println("phone value: ", phone)
	ap := AndroidPhone{
		Name: "AndroidPhone",
	}

	fmt.Println("--------------------androidPhone-------------------------")

	var androidPhone Phone
	androidPhone = ap

	androidPhone.Call()
	androidPhone.SendMsg("android")

	fmt.Println("--------------------gameConsole-------------------------")
	var game gameConsole
	game = ap

	game.Play()

	fmt.Println("------------------------ALLInOne------------------------")
	var all AllInOne
	all = ap
	all.Play()
	all.SendMsg("ok")
	all.Call()

}
