package repository

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"uala/go-workshop/pkg/dto"
)

const (
	TableName = "Contacts"
)

type Repository interface {
	Setup()
	Insert(contact dto.Contact) (dto.Contact, error)
}

type LambdaRepository struct {
	TableName string
	svc       *dynamodb.DynamoDB
}

func (r *LambdaRepository) Setup() {
	r.TableName = TableName

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	r.svc = dynamodb.New(sess)
}

func (r *LambdaRepository) Insert(contact dto.Contact) (dto.Contact, error) {
	// Convert the Record Go type to dynamodb attribute value type using MarshalMap
	item, err := dynamodbattribute.MarshalMap(contact)
	if err != nil {
		return dto.Contact{}, err
	}

	// Declare a new PutItemInput
	input := &dynamodb.PutItemInput{
		Item: item,
		TableName: aws.String(r.TableName),
	}

	// Put new item into the dynamodb table
	_, err = r.svc.PutItem(input)
	if err != nil {
		return dto.Contact{}, err
	}

	return contact, nil
}
