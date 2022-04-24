package main

type ChainLogger interface {
	Next(s string)
}

type ClosureChain struct {
	NextChain ChainLogger
	Closure func(string)
}

func (c *ClosureChain) Next(s string) {
	if c.Closure != nil {
		c.Closure(s)
	}
	if c.NextChain != nil {
		c.Next(s)
	}
}
