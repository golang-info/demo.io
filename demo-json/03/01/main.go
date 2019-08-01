package main

import (
	"encoding/json"
	"os"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	//如果Child字段为nil 编码JSON可忽略
	Child *Person `json:"child,omitempty"`
}

func main() {
	person := Person{
		Name: "John",
		Age:  40,
		Child: &Person{
			Name: "Tom",
			Age:  20,
		},
	}

	//File 类型实现了io.Writer接口
	file, _ := os.Create("./person.json")

	//根据io.Writer创建encoder 然后调用Encode()方法将对象编码成JSON
	json.NewEncoder(file).Encode(&person)

}
