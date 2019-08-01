//  https://www.cnblogs.com/liuhe688/p/10971327.html
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type (
	person struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	result struct {
		Group   string   `json:"group"`
		Persons []person `json:"persons"`
	}
)

func main() {
	var data result

	bytes, _ := ioutil.ReadFile("/Users/liuyanchao/Documents/11-goworking/src/demo.io/demo-json/01/data.json")
	fmt.Println("**** data.json content: ***")
	fmt.Println(string(bytes))

	json.Unmarshal(bytes, &data)
	fmt.Println("*** unmarshal result: ***")
	fmt.Println(data)

	data.Group = "enginner"
	//newBytes, _ := json.Marshal(&data)
	newBytes, _ := json.MarshalIndent(&data, "", "    ")
	fmt.Println("*** updata content: ***")
	fmt.Println(string(newBytes))
	ioutil.WriteFile("/Users/liuyanchao/Documents/11-goworking/src/demo.io/demo-json/01/data_new.json", newBytes, os.ModeAppend)

}
