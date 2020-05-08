package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main() {
	s := "bash /etc/profile"

	//strings -> []strings
	c := func (str string) []string {
		var d []string
		for _, v := range strings.Fields(str) {
			d = append(d, v)
		}
		return d
	}(s)

	fmt.Println("c: ",c)
	fmt.Println("c type: ", reflect.TypeOf(c))
}