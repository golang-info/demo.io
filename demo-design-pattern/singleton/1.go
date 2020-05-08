package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	lock *sync.Mutex = &sync.Mutex{}
	instance *Singleton
)

type Singleton struct {

}

func GetInstance() *Singleton {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()
		if instance == nil {
			instance = &Singleton{}
			fmt.Println("instance...")
		}
	}
	return instance
}

func main() {
	var s *Singleton
	s = GetInstance()
	time.Sleep(2 * time.Second)
	s = GetInstance()
	fmt.Println(s)
}
