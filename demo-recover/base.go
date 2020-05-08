package main

import (
	"fmt"
)

func main() {
	defer func() {
		fmt.Println("1")
	}()

	defer func(){
		if err :=  recover(); err != nil {
			fmt.Println(err)
		}
	}()
	panic("fault")
	fmt.Println("2")
}