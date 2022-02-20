package models

type Contact struct {
	Uuid      string `dynamodbav:"Uuid" json:"id"`
	FirstName string `dynamodbav:"FirstName" json:"first_name"`
	LastName  string `dynamodbav:"LastName" json:"last_name"`
	Status    string `dynamodbav:"Status" json:"status"`
}
