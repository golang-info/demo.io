package main

import "fmt"
import "time"

func main() {

	p := fmt.Println
	fmt.Println("Hello World!")
	fmt.Println("go" + "lang")
	fmt.Println("1+1 =", 1+1)
	fmt.Println("7.0/3.0=", 7.0/3.0)

	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)

	fmt.Println("ok ")
	fmt.Println("test")

	now := time.Now()
	fmt.Println(now)

	tmp := time.Date(2009, 11, 17, 20, 34, 58, 651387237, time.UTC)
	fmt.Println(tmp)

	p(tmp.Year())

	fmt.Println("ok")
}
