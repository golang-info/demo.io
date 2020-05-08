package main

import "fmt"

type CPU interface {
	Calculate()
}

type Memory interface {
	Storage()
}

type Card interface {
	Display()
}

type Computer struct {
	CPU
	Memory
	Card
}

func NewComputer(cpu CPU, memory Memory, card Card) *Computer {
	return &Computer{
		CPU:    cpu,
		Memory: memory,
		Card:   card,
	}
}

func (c *Computer) DoWork() {
	c.CPU.Calculate()
	c.Memory.Storage()
	c.Card.Display()
}

type IntelCpu struct {
	Name string
}

func (i *IntelCpu) Calculate() {
	fmt.Println(	i.Name , " 进行了计算")
}

type IntelMemory struct {
	Name string
}

func (i *IntelMemory) Storage() {
	fmt.Println(i.Name, " 进行了存储")
}

type IntelCard struct {
	Name string
}

func (i *IntelCard) Display() {
	fmt.Println(i.Name, " 进行了显示")
}

type OtherCpu struct {
	Name string
}

func (ocpu *OtherCpu) Calculate() {
	fmt.Println(ocpu.Name, " 进行了计算")
}

type OtherMemory struct {
	Name string
}

func (om *OtherMemory) Storage() {
	fmt.Println(om.Name, " 进行了存储")
}

type OtherCard struct {
	Name string
}

func (oc *OtherCard) Display() {
	fmt.Println(oc.Name, " 进行了显示")
}


func main() {
	computer1 := NewComputer(&IntelCpu{Name: "ICPU"}, &IntelMemory{"IMEM"}, &IntelCard{Name: "ICARD"})
	computer2 := NewComputer(&OtherCpu{Name: "OCPU"}, &OtherMemory{Name: "OMEM"}, &OtherCard{Name: "OCARD"})

	computer1.DoWork()
	computer2.DoWork()
}


