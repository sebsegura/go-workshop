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
	Insert = "INSERT"
)

type Repository interface {
	Insert(contact dto.Contact) (dto.Contact, error)
}

type LambdaRepository struct {
	TableName string
	svc       *dynamodb.DynamoDB
}

func New() Repository {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	return &LambdaRepository{
		TableName: TableName,
		svc:       dynamodb.New(sess),
	}
}

func (r *LambdaRepository) Insert(contact dto.Contact) (dto.Contact, error) {
	// Convert the Record Go type to dynamodb attribute value type using MarshalMap
	item, err := dynamodbattribute.MarshalMap(contact)
	if err != nil {
		return dto.Contact{}, &dto.DynamoDbError{
			Op: Insert,
			Err: dto.InvalidInputError,
		}
	}

	// Declare a new PutItemInput
	input := &dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String(r.TableName),
	}

	// Put new item into the dynamodb table
	_, err = r.svc.PutItem(input)
	if err != nil {
		return dto.Contact{}, &dto.DynamoDbError{
			Op: Insert,
			Err: dto.InsertionError,
		}
	}

	return contact, nil
}
