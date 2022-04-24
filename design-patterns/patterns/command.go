package main

import "fmt"

type Command interface {
	Execute()
}

type Console struct {
	message string
}

func (c *Console) Execute() {
	fmt.Println(c.message)
}

func CreateCommand(s string) Command {
	return &Console{
		message: s,
	}
}

type CommandQueue struct {
	queue []Command
}

func (p *CommandQueue) AddCommand(c Command) {
	p.queue = append(p.queue, c)
	if len(p.queue) == 3 {
		for _, cmd := range p.queue {
			cmd.Execute()
		}
		p.queue = make([]Command, 3)
	}
}

func main() {
	queue := CommandQueue{}

	queue.AddCommand(CreateCommand("first"))
	queue.AddCommand(CreateCommand("second"))
	queue.AddCommand(CreateCommand("third"))
}