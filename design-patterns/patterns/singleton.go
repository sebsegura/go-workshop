package main

import (
	"fmt"
	"sync"
)

type singleton struct{}

var (
	lock     = &sync.Mutex{}
	instance *singleton
)

func getInstance() *singleton {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()
		if instance == nil {
			fmt.Println("Creating single instance...")
			instance = &singleton{}
		} else {
			fmt.Println("Single instance already created!")
		}
	} else {
		fmt.Println("Single instance already created!")
	}

	return instance
}

func main() {
	for i := 0; i < 4; i++ {
		getInstance()
	}
}
