package main

import "fmt"

type Weapon int

const (
	Arrow Weapon = iota   //开始生成枚举值， 默认为0
	Shuriken
	SniperRifle
	Rifle
	Blower
)

func main() {
	//输出所有枚举值
	fmt.Println(Arrow, Shuriken, SniperRifle, Rifle, Blower)

	// 使用枚举类型并赋初值
	var weapon Weapon = Blower
	fmt.Println(weapon)

}
