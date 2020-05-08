package main

import "fmt"

func main() {
	var s = make([]string, 0)
	var m = make(map[string]string)

	s = append(s, "some thing")
	m["some key"] = "some value"

	fmt.Println(s, m)
	fmt.Println(len(s), cap(s))
	fmt.Println(len(m))

	//------- pointers -----------
	var count = int(42)
	ptr := &count
	fmt.Println("&ptr: ", &ptr, "&count: ", &count)
	fmt.Println("ptr: ", ptr)
	fmt.Println(*ptr)

	*ptr = 100
	fmt.Println(count)


}
