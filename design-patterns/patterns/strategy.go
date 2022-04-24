package main

import (
	"flag"
	"log"
)

type PrintStrategy interface {
	Print() error
}

type ConsoleSquare struct {}

type ImageSquare struct {
	DestinationFilePath string
}

type TextSquare struct {}

func (c *ConsoleSquare) Print() error {
	return nil
}

func (t *ImageSquare) Print() error {
	return nil
}

func (t *TextSquare) Print() error {
	return nil
}

var output = flag.String("output", "console", "Input")

func main() {
	flag.Parse()
	var activeStrategy PrintStrategy

	switch *output {
	case "console":
		activeStrategy = &ConsoleSquare{}
	case "image":
		activeStrategy = &ImageSquare{"path"}
	default:
		activeStrategy = &TextSquare{}
	}

	if err := activeStrategy.Print(); err != nil {
		log.Fatal(err)
	}
}
