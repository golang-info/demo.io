package main

import "fmt"

func main() {
	fmt.Print("\x1b[4;30;44m") //设置颜色样式
	fmt.Print("Hello World")  //打印文本内容
	fmt.Println("\x1b[0m")
} 
