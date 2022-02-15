package common

import (
	ddb "github.com/seb7887/aws-lib/dynamodb"
	"uala/go-workshop/pkg/dto"
)

var (
	MockRequest = dto.Request{
		FirstName: "john",
		LastName:  "doe",
	}
	MockInternalError *ddb.InternalError
	MockLambdaError   *dto.LambdaError
	MockContact       = ddb.Contact{
		ID:        "",
		FirstName: MockRequest.FirstName,
		LastName:  MockRequest.LastName,
		Status:    "CREATED",
	}
)
