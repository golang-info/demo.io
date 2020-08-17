package main

/*
	ref:
	https://learnku.com/docs/build-web-application-with-golang/096-encryption-and-decryption-of-data/3214
*/

import (
	"encoding/base64"
	"fmt"
)

func main() {

	// 函数原型 解密
	// base64.StdEncoding.DecodeString(s string) ([]byte, error)
	// 函数原型 加密
	// base64.StdEncoding.EncodeToString(src []byte) string

	// 演示base64编码
	str := "I am ok"
	encodeString := base64.StdEncoding.EncodeToString([]byte(str))
	fmt.Println(encodeString)

	// 对上面的编码结果进行base64解码
	decodeBytes, err := base64.StdEncoding.DecodeString(encodeString)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(decodeBytes))

}
