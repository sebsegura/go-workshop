package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"log"
)

type handler struct {}

func (h *handler) Create() {
	// do some stuff
}

func main() {
	h := handler{}
	log.Println("Initializing handler...")
	lambda.Start(h.Create)
}