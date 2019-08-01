package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	//如果Child字段为nil 编码JSON可忽略
	Child *Person `json:"child, omitempty"`
}

func main() {
	//------------------------
	var person Person

	//File 类型也实现了io.Reader接口
	file, _ := os.Open("./person.json")

	//根据io.Reader创建Decoder 然后调用Decode()方法将JSON解析成对象
	json.NewDecoder(file).Decode(&person)

	fmt.Println(person)
	fmt.Println(*person.Child)

}
