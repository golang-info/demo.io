package main

import (
	"fmt"
	"os"
	"runtime/trace"
)

// trace 的编程过程
// 1 创建文件
// 2 启动
// 3 停止

// go run trace.go
// go tool trace trace.out
// 2020/08/07 15:44:08 Parsing trace...
// 2020/08/07 15:44:08 Splitting trace...
// 2020/08/07 15:44:08 Opening browser. Trace viewer is listening on http://127.0.0.1:54346

func main() {
	// 1 创建文件
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	// 2 启动
	err = trace.Start(f)
	if err != nil {
		panic(err)
	}

	// 业务逻辑代码
	fmt.Println("Hello GMP")

	// 3 停止
	trace.Stop()
}
