package main

import "fmt"

func sample() {
	fmt.Println("Inside the sample()")
}

func display(a int) {
	fmt.Println(a)
}

func main() {
	//sample() will be invoked only after executing the statements of main()
	defer sample()
	//sample()
	defer display(1)
	defer display(2)
	defer display(3)
	fmt.Println("Inside the main()")
}
