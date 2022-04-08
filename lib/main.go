package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/sns"
	"sync"
)

type Order struct {
	OrderId string `dynamodbav:"OrderId" json:"OrderId"`
}

type OrderBuilder interface {
	SetOrderId(id string) OrderBuilder
	Build() Order
}

type orderBuilder struct {
	id string
}

func NewOrderBuilder() OrderBuilder {
	return &orderBuilder{}
}

func (b *orderBuilder) SetOrderId(id string) OrderBuilder {
	b.id = id
	return b
}

func (b *orderBuilder) Build() Order {
	return Order{OrderId: b.id}
}

func (o *Order) ToJsonStr() string {
	b, err := json.Marshal(o)
	if err != nil {
		return ""
	}
	return string(b)
}

type ErrorMessage struct {
	Message string `json:"errorMessage"`
}

func (e *ErrorMessage) Error() string {
	return e.Message
}

type ErrorMsg struct {
	HttpStatus string `json:"httpStatus"`
	Message    string `json:"msg"`
	Code       string `json:"code"`
}

/***********************
CODES AND DESCRIPTIONS
************************/
const (
	GenericError    = 999
	InputValidation = 1000 + iota
	WrongClientId
	InvalidAmountFormat
	RegisterNotFound
	MarshalError
	UnmarshalError
	UpdatingError
	InsertionError
)

var ErrorMessages = map[int]string{
	GenericError:        "Error executing action.",
	InputValidation:     "Input validation FAILED",
	WrongClientId:       "Wrong clientId",
	InvalidAmountFormat: "Amount validation FAILED. Amount must have only numbers and a single point",
	RegisterNotFound: "No register found",
	MarshalError: "Marshal Error",
	UnmarshalError: "Unmarshal error",
	UpdatingError: "Error updating Item",
	InsertionError: "Error in the insertion of the item",
}

var ErrorsDictionary = map[int]ErrorMsg{
	InputValidation: {
		HttpStatus: "400",
		Message:    ErrorMessages[InputValidation],
		Code:       fmt.Sprintf("%d", InputValidation),
	},
	WrongClientId: {
		HttpStatus: "400",
		Message:    ErrorMessages[WrongClientId],
		Code:       fmt.Sprintf("%d", WrongClientId),
	},
	InvalidAmountFormat: {
		HttpStatus: "400",
		Message:    ErrorMessages[InvalidAmountFormat],
		Code:       fmt.Sprintf("%d", InvalidAmountFormat),
	},
	GenericError: {
		HttpStatus: "400",
		Message:    ErrorMessages[GenericError],
		Code:       fmt.Sprintf("%d", GenericError),
	},
	RegisterNotFound: {
		HttpStatus: "400",
		Message:    ErrorMessages[RegisterNotFound],
		Code:       fmt.Sprintf("%d", RegisterNotFound),
	},
	UnmarshalError: {
		HttpStatus: "400",
		Message:    ErrorMessages[UnmarshalError],
		Code:       fmt.Sprintf("%d", UnmarshalError),
	},
	UpdatingError: {
		HttpStatus: "400",
		Message:    ErrorMessages[UpdatingError],
		Code:       fmt.Sprintf("%d", UpdatingError),
	},
	InsertionError: {
		HttpStatus: "400",
		Message:    ErrorMessages[InsertionError],
		Code:       fmt.Sprintf("%d", InsertionError),
	},
	MarshalError: {
		HttpStatus: "400",
		Message:    ErrorMessages[MarshalError],
		Code:       fmt.Sprintf("%d", MarshalError),
	},
}

func (e *ErrorMsg) ToJson() string {
	b, err := json.Marshal(e)
	if err != nil {
		return ""
	}
	return string(b)
}

func (e *ErrorMsg) ToErrorMessage() *ErrorMessage {

	return &ErrorMessage{Message: e.ToJson()}
}

var (
	PutItemError = errors.New("PutItem error")
	SnsError = errors.New("SNS error")
)

type Repository interface {
	CreateOrder(order Order) error
}

var (
	db *dynamodb.DynamoDB
	once sync.Once
)

func GetDbConn() *dynamodb.DynamoDB {
	once.Do(func() {
		sess := session.Must(session.NewSessionWithOptions(session.Options{
			SharedConfigState: session.SharedConfigEnable,
		}))
		db = dynamodb.New(sess)
	})
	return db
}

func CreateOrder(order Order) error {
	ddb := GetDbConn()

	item, err := dynamodbattribute.MarshalMap(order)
	if err != nil {
		return err
	}

	input := &dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String("Orders"),
	}

	_, err = ddb.PutItem(input)
	return err
}

var (
	snsSvc *sns.SNS
)

func GetSnsConn() *sns.SNS {
	once.Do(func() {
		snsSess := session.Must(session.NewSession(&aws.Config{Region: aws.String("us-east-1")}))
		snsSvc = sns.New(snsSess)
	})
	return snsSvc
}

func PublishOrder(order Order) error {
	svc := GetSnsConn()
	_, err := svc.Publish(&sns.PublishInput{
		Message:  aws.String(order.ToJsonStr()),
		TopicArn: aws.String("topicArn"),
	})
	return err
}

func Processor(order Order) error {
	err := CreateOrder(order)
	if err != nil {
		return PutItemError
	}

	err = PublishOrder(order)
	if err != nil {
		return SnsError
	}

	return nil
}

func Handler() error {
	builder := NewOrderBuilder()
	order := builder.SetOrderId("1").Build()

	err := Processor(order)
	if err != nil {
		// si es error de tipo PutItemError??
		if errors.Is(err, PutItemError) {
			errMsg := ErrorsDictionary[InsertionError]
			return errMsg.ToErrorMessage()
		}
		// si es otro tipo de error
		errMsg := ErrorsDictionary[GenericError]
		return errMsg.ToErrorMessage()
	}

	return nil
}

func main() {
	_ = Handler()
}