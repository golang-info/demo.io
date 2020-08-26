package main

import "fmt"

func main() {
	frontend := make(chan string)
	backend := make(chan string)
	quit := make(chan string)

	go send(frontend, backend, quit)
	receive(frontend, backend, quit)

}

func send(f, b, q chan<- string) {
	data := []string{"React", "NodeJS", "Vue", "Flask", "Angular", "Laravel"}
	for i := 0; i < len(data); i++ {
		if i%2 == 0 {
			f <- data[i]
		} else {
			b <- data[i]
		}
	}

	q <- "finished"
}

func receive(f, b, q <-chan string) {
	for {
		select {
		case v := <-f:
			fmt.Println("Front end Dev:", v)
		case v := <-b:
			fmt.Println("Back End Dev:", v)
		case v := <-q:
			fmt.Println("This program is", v)
			return
		}
	}
}
