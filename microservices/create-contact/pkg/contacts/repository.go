package contacts

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"seb7887/create-contact/internal/storage"
	"seb7887/create-contact/pkg/models"
)

const (
	ContactsTable = "Contacts"
)

type ContactRepository interface {
	PutItem(contact models.Contact) error
}

type contactsRepository struct {
	table string
	ddb   *dynamodb.DynamoDB
}

func NewContactsRepository() ContactRepository {
	ddb := storage.GetDb()

	return &contactsRepository{
		table: ContactsTable,
		ddb:   ddb,
	}
}

func (r *contactsRepository) PutItem(contact models.Contact) error {
	item, err := dynamodbattribute.MarshalMap(contact)
	if err != nil {
		return &models.MarshalError{
			In:  contact,
			Err: err,
		}
	}

	input := &dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String(r.table),
	}

	_, err = r.ddb.PutItem(input)
	if err != nil {
		fmt.Println(err)
		return &models.InternalError{
			Op:  "PutItem",
			Err: err,
		}
	}

	return nil
}
