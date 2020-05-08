package main

func main() {
	go println("goroutine message")
	println("main function message")
}
