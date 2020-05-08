package main

import "fmt"

func fub(i int) int {
	if i == 0 || i == 1 {
		return 1
	} else {
		return fub(i - 1) + fub(i - 2)
	}
}

func main() {
	sum := fub(10)
	fmt.Println("sum: ", sum)
	fmt.Println("done")
}
