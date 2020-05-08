package main

import (
	"fmt"
)

//---------------抽象层--------------
type Card interface {
	Display()
}

type Memory interface {
	Storage()
}

type CPU interface {
	Calculate()
}

type Computer struct {
	cpu CPU
	memory Memory
	card Card
}

func NewComputer(cpu CPU, mem Memory, card Card) *Computer {
	return &Computer{
		cpu:    cpu,
		memory: mem,
		card:   card,
	}
}

func (this *Computer) DoWork() {
	this.cpu.Calculate()
	this.memory.Storage()
	this.card.Display()
}

//---------------实现层----------------
//intel
type IntelCPU struct {
	CPU
}

func (this *IntelCPU) Calculate() {
	fmt.Println("Intel CPU 开始计算了。。。")
}

type IntelMemory struct {
	Memory
}

func (this *IntelMemory) Storage() {
	fmt.Println("Intel Memory 开始存储了。。。")
}

type IntelCard struct {
	Card
}

func (this *IntelCard) Display() {
	fmt.Println("Intel Card 开始显示了。。。")
}

//Kingston
type KinstonMemory struct {
	Memory
}

func (this *KinstonMemory) Storage() {
	fmt.Println("Kongstaon memory storage...")
}

//nmidia
type NvidiaCard struct {
	Card
}

func (this *NvidiaCard) Display() {
	fmt.Println("Nvidia card display...")
}

//------------业务逻辑层--------------
func main() {
	//intel 系列的电脑
	com1 := NewComputer(&IntelCPU{}, &IntelMemory{}, &IntelCard{})
	com1.DoWork()

	//杂牌子
	com2 := NewComputer(&IntelCPU{}, &KinstonMemory{}, &NvidiaCard{})
	com2.DoWork()
}