package main

import (
	"encoding/json"
	"fmt"
)

/* json 解析非结构化数据的方法
每个字符串对应一个 JSON 属性的名称， interface{} 类型对应它的值，这个值可以是任意类型。
在代码中，其数据类型通过对 interface{} 进行断言就可以获取到。
而且这些映射解析动作可以迭代执行，如此一来，可变数量的键通过一次循环中就能够处理完成。 */

func main() {
	birdJson := `{
		"birds": {
		  "pigeon": "likes to perch on rocks",
		  "eagle": "bird of prey"
		},
		"animals": "none"
	  }`
	// var result map[string]interface{}
	result := make(map[string]interface{})
	//result := new(map[string]interface{})
	json.Unmarshal([]byte(birdJson), &result)

	// The object stored in the "birds" key is alse stored as
	// a map[string]interface{} type, and its type is asserted from
	// the interface{} type
	birds := result["birds"].(map[string]interface{})

	for key, value := range birds {
		// Echo value is an interface{} type, that is type asserted as a string
		fmt.Println(key, value.(string))
	}
}
