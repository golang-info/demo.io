package main

/*
	https://blog.csdn.net/skh2015java/article/details/102748222
*/

import (
	"net/http"
	_ "net/http/pprof"
)
import "fmt"

func main() {
	fmt.Println("Hello world!")
	go func() {
		http.ListenAndServe("0.0.0.0:6060", nil)
	}()

	select {}
}
