package main

import "fmt"

type Stu struct {
	Name string
}

func main() {
	fmt.Printf("%v\n", Stu{"Tom"})
	fmt.Printf("%+v\n", Stu{"Tom"})
}