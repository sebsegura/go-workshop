package storage

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"seb7887/create-contact/internal/config"
	"sync"
)

const (
	REGION = "us-west-2"
)

var (
	ddb  *dynamodb.DynamoDB
	once sync.Once
)

func GetDb() *dynamodb.DynamoDB {
	once.Do(func() {
		sess, err := session.NewSession(&aws.Config{
			Region:      aws.String(REGION),
			Endpoint:    aws.String(config.GetConfig().Endpoint),
			Credentials: credentials.NewEnvCredentials(),
		})
		if err != nil {
			panic("Cannot connect with database")
		}
		ddb = dynamodb.New(sess)
	})
	return ddb
}
