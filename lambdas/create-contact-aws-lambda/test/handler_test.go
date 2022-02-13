package test

import (
	"context"
	"errors"
	"github.com/aws/aws-lambda-go/lambdacontext"
	ddb "github.com/seb7887/aws-lib/dynamodb"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"testing"
	"uala/go-workshop/pkg/dto"
	"uala/go-workshop/pkg/handler"
	"uala/go-workshop/test/mocks"
)

var (
	ct = &lambdacontext.LambdaContext{
		AwsRequestID:       "awsRequestId1234",
		InvokedFunctionArn: "arn:aws:lambda:xxx",
		Identity:           lambdacontext.CognitoIdentity{},
		ClientContext:      lambdacontext.ClientContext{},
	}
	ctx = lambdacontext.NewContext(context.TODO(), ct)
)

func TestHandler_Ok(t *testing.T) {
	var (
		mockProcessor = mocks.MockProcessor{}
	)

	mockProcessor.On("Process", mockRequest).Return(mockContact, nil)

	h := handler.New(&mockProcessor)

	res, err := h.Create(ctx, mockRequest)

	// Assertions
	require.NoError(t, err, "should not return an error")
	assert.Equal(t, http.StatusOK, res.StatusCode)
}

func TestHandler_ValidationError(t *testing.T) {
	var (
		mockProcessor = mocks.MockProcessor{}
		badRequest    = dto.Request{
			FirstName: "",
			LastName:  "",
		}
	)

	h := handler.New(&mockProcessor)

	res, err := h.Create(ctx, badRequest)

	// Assertions
	require.NoError(t, err, "should not return an error")
	assert.Equal(t, http.StatusBadRequest, res.StatusCode)
}

func TestHandler_ProcessorError(t *testing.T) {
	var (
		mockProcessor = mocks.MockProcessor{}
	)

	mockProcessor.On("Process", mockRequest).Return(ddb.Contact{}, &ddb.InternalError{
		Op:  "PutItem",
		Err: errors.New("something bad happened"),
	})

	h := handler.New(&mockProcessor)

	res, err := h.Create(ctx, mockRequest)

	// Assertions
	require.NoError(t, err, "should not return an error")
	assert.Equal(t, http.StatusInternalServerError, res.StatusCode)
}
