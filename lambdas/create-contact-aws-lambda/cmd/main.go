package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	ddb "github.com/seb7887/aws-lib/dynamodb"
	"uala/go-workshop/internal/processor"
	"uala/go-workshop/pkg/handler"
)

// Punto de entrada de cualquier programa hecho de Go
func main() {
	r := ddb.NewContactsRepository()
	p := processor.New(r)
	h := handler.New(p)

	lambda.Start(h.Create)
}
