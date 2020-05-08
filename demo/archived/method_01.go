package main

import (
	"fmt"
)

type TZ int

func (tz *TZ) Increase(num int) {
	*tz += TZ(num)
}

func main() {
	var a TZ
	a = 100
	a.Increase(100)
	fmt.Println(a)
}