package main

import (
	"fmt"
	"sync"
)

type single struct {
	count int
}

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

func (s *single) AddOne() int {
	s.count++
	return s.count
}

func main() {
	counter := getSingleInstance()
	fmt.Println(counter.AddOne())
}
