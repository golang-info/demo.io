package main
/*
	https://www.cnblogs.com/dajianshi/p/4235142.html
*/
import (
	"fmt"
)

func main() {
	fmt.Println("ok")
	var ss []string
	fmt.Printf("[ local print ]\t:\t length: %v\taddr:%p\tisnil:%v\n", len(ss), ss, ss==nil)

	for i:=0;i<10;i++{
		ss=append(ss, fmt.Sprintf("s%d", i))
	}
	fmt.Printf("[ local print ]\t:\t length: %v\taddr:%p\tisnil:%v\n", len(ss), ss, ss==nil)
	fmt.Println("after append: ", ss)

	index := 0
	near := append([]string{}, ss[index:]...)
	ss = append(ss[0:index], "interted")
	ss = append(ss, near...)
	fmt.Println("after interted: ", ss)

	if ss[0] == "interted" {
		ss[0] = "newInterted"
		fmt.Println("after update index 0 interted: ", ss)
	}
}