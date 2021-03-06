package main

import "fmt"
import "sync"

func main() {
	var scene sync.Map

	scene.Store("greece", 97)
	scene.Store("london", 100)
	scene.Store("egype", 200)

	fmt.Println(scene.Load("london"))

	scene.Delete("london")

	scene.Range(func(k, v interface{}) bool {
		fmt.Println("iterate: ", k, v)
		return true
	})
}
