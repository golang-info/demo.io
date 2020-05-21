package main

import (
	"fmt"
	messagebus "github.com/vardius/message-bus"
	"sync"
)

func main() {
	queueSize := 10
	bus := messagebus.New(queueSize)

	var wg sync.WaitGroup
	wg.Add(2)

	_ = bus.Subscribe("topic", func(v bool) {
		defer wg.Done()
		fmt.Println(v)
	})

	_ = bus.Subscribe("topic", func(v bool) {
		defer wg.Done()
		fmt.Println(v)
	})
	fmt.Println("ok")

	bus.Publish("topic", true)
	wg.Wait()
}
