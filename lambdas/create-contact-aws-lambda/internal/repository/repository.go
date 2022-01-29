package repository

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"uala/go-workshop/pkg/dto"
)

type Repository interface {
	Setup()
	Insert(contact dto.Contact) (dto.Contact, error)
}

type LambdaRepository struct {
	TableName string
	svc       *dynamodb.DynamoDB
}

// TODO: Setup -> instanciar cliente dynamodb usando credentials en shared config
func (r *LambdaRepository) Setup() {}

// TODO: Insert -> insertar un nuevo elemento en la tabla de dynamodb
func (r *LambdaRepository) Insert(contact dto.Contact) (dto.Contact, error) {
	// Convert the Record Go type to dynamodb attribute value type using MarshalMap

	// Declare a new PutItemInput

	// Put new item into the dynamodb table

	return contact, nil
}
