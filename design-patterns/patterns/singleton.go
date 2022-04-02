package main

import "fmt"

type singleton struct {
	count int
}

var (
	instance *singleton
)

func getInstance() *singleton {
	if instance == nil {
		instance = new(singleton)
	}
	return instance
}

func (s *singleton) AddOne() int {
	s.count++
	return s.count
}

func main() {
	counter := getInstance()
	fmt.Println(counter.AddOne())
}
