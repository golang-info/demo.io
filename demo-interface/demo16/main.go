package main

/*
http://c.biancheng.net/view/78.html
*/

import (
	"fmt"
)

//定义一个数据写入器
type DataWriter interface {
	WriteData(data interface{}) error
}

//定义文件结构，用于实现DataWriter
type file struct {

}

//实现DataWriter接口的WriterData方法
func (d *file) WriteData(data interface{}) error {
	//模拟写入数据
	fmt.Println("WriteData: ", data)
	return nil
}

func main() {
	//实例化file
	f := new(file)

	//声明一个DataWriter的接口
	var writer DataWriter

	//将接口赋值f， 也就是*file类型
	//将*file类型的f赋值给DataWriter接口的writer，虽然两个变量类型不一致， 但是writer是一个接口，
	// 且f已经完全实现了DataWriter接口的所有方法，因此赋值是成功的
	writer = f

	//使用DataWriter接口进行数据写入
	writer.WriteData("dataA")
}
