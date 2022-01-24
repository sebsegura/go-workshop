package repository

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"seb7887/create-contact/pkg/dto"
)

type Repository interface {
	Setup()
	Insert(contact dto.Contact) (dto.Contact, error)
}

type ContactRepository struct {
	TableName string
	svc *dynamodb.DynamoDB
}

func (r *ContactRepository) Setup() {
	r.TableName = "Contacts"
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	r.svc = dynamodb.New(sess)
}

func (r *ContactRepository) Insert(contact dto.Contact) (dto.Contact, error) {
	item, err := dynamodbattribute.MarshalMap(contact)
	if err != nil {
		return dto.Contact{}, err
	}

	input := &dynamodb.PutItemInput{
		Item: item,
		TableName: aws.String(r.TableName),
	}

	_, err = r.svc.PutItem(input)
	if err != nil {
		return dto.Contact{}, err
	}

	return contact, nil
}
