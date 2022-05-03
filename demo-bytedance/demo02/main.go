package main

func main() {
	ch := make(chan int, 0)

	ch <- 666
	go func() {
		<- ch
	}()
}
