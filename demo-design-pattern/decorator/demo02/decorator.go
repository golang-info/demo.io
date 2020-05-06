package main

import "fmt"

//定义公用接口
type Shower interface {
	Show()
}

//实现了公用接口的具体的类
type Person struct {
	Name string
}

func (p *Person) Show() {
	fmt.Println("装扮的"， p.Name)
}

//对Person进行装饰， show的前后增加额外操作
type Finery struct {
	Shower
}

func (self *Finery) Decorate(component Shower) {
	self.Shower = component
} 

func (self *Finery) Show() {
	if self.Shower == nil {
		return
	}
	fmt.Println("准备起床")
	self.Shower.Show()
	fmt.Println("穿上漂亮的衣服")
}

//对Person进行装饰， show之后增加额外操作
type TShirts struct {
	Shower
}

func (self *TShirts) Decorate(component Shower) {
	self.Shower = component
}

func (self *TShirts) Show() {
	if self.Shower == nil {
		return
	}
	self.Shower.Show()
	fmt.Println("穿上大T恤")
}

//对Person进行装饰， show的之后增加额外操作
type Sneakers struct {
	Shower
}

func (self *Sneakers) Decorate(component Shower) {
	self.Shower = component
}

func (self *Sneakers) Show() {
	if self.Shower == nil {
		return
	}
	if self.Shower.Show()
	fmt.Println("穿上破球鞋")
}
