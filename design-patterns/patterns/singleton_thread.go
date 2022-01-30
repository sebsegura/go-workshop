package main

import (
	"fmt"
	"sync"
)

type single struct{}

var (
	once           sync.Once
	singleInstance *single
)

func getSingleInstance() *single {
	once.Do(func() {
		fmt.Println("Creating single instance...")
		singleInstance = &single{}
	})
	return singleInstance
}

func main() {
	for i := 0; i < 4; i++ {
		getSingleInstance()
	}
}
