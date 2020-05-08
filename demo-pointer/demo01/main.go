package main

import "fmt"

func main() {
	var a int 
	var p *int

	fmt.Printf("a type is %T \n",a)
	fmt.Printf("p type is %T \n",p)
	fmt.Printf("指针变量 p 的值为： %x\n",p)
	if (p == nil) {
		fmt.Printf("指针变量 p 的值为： %x\n",p)
	}

	a = 100
	p = &a

	fmt.Println(a)
	fmt.Println(&a)
	fmt.Println(p)
	fmt.Println(*p)

	fmt.Println("------------------------------------")

	*p = 200

	fmt.Println(a)
	fmt.Println(&a)
	fmt.Println(p)
	fmt.Println(*p)
	fmt.Println(&p)



}

