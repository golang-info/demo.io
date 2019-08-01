package main

import (
	"encoding/json"
	"fmt"
)

var JSON = `{
	"group": "programmer",
	"persons": [
		{
			"name": "Tom",
			"age": 40
		},
		{
			"name": "Jhon",
			"age": 20
		}
	]
}`

func main() {
	var bytes []byte
	var data map[string]interface{}

	//将字符串转为字节切片
	bytes = []byte(JSON)

	//将字节切片映射到map上
	json.Unmarshal(bytes, &data)

	fmt.Println("group: ", data["group"])

	//转为[]interface{}类型
	persons := data["persons"].([]interface{})

	for index, item := range persons {
		//类型转换
		person := item.(map[string]interface{})

		age := person["age"]

		//更改年龄
		person["age"] = age.(float64) + 1

		//打印最新个人信息
		fmt.Println("person", index, ":", person["name"].(string), age, "->", person["age"].(float64))
	}

	//序列化为JSON字符串
	newBytes, _ := json.MarshalIndent(&data, "", "    ")

	//打印新的JSON数据
	fmt.Println(string(newBytes))
}
