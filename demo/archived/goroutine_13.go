package main

import "fmt"

func main() {
	ch := make(chan string)

	go func(){
		ch <- "hello01"
		ch <- "hello02"
		ch <- "hello03"
		ch <- "hello04"
		close(ch)
	}()


	for res := range ch {
		fmt.Println(res)
	}
}