package test

import (
	ddb "github.com/seb7887/aws-lib/dynamodb"
	"uala/go-workshop/pkg/dto"
)

var (
	mockRequest = dto.Request{
		FirstName: "john",
		LastName:  "doe",
	}
	mockInternalError *ddb.InternalError
	mockLambdaError   *dto.LambdaError
	mockContact       = ddb.Contact{
		ID:        "",
		FirstName: mockRequest.FirstName,
		LastName:  mockRequest.LastName,
		Status:    "CREATED",
	}
)
