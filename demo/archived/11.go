package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("cpus: ", runtime.NumCPU())
	fmt.Println("goroot: ", runtime.GOROOT())
	fmt.Println("archive: ", runtime.GOOS)
	fmt.Println(runtime.Version(), runtime.GOARCH, runtime.MemStats{})
	//n := runtime.Stack([]int{500,) true)
}
