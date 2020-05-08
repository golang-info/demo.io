package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	once sync.Once
	instance *Singleton
)

type Singleton struct {

}

func GetInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{}
		fmt.Println("instance...")
	})
	return instance
}

func main() {
	var s *Singleton
	s = GetInstance()
	time.Sleep(1 * time.Second)
	s = GetInstance()
	fmt.Println(s)
}
