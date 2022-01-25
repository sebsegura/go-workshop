package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"uala/go-workshop/pkg/handler"
)

// Punto de entrada de cualquier programa hecho de Go
func main() {
	h := handler.LambdaHandler{}

	lambda.Start(h.Create)
}