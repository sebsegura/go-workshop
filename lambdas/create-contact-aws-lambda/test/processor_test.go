package test

import (
	"errors"
	ddb "github.com/seb7887/aws-lib/dynamodb"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"uala/go-workshop/internal/processor"
	"uala/go-workshop/test/mocks"
)

func TestProcessor_Process_Ok(t *testing.T) {
	var (
		mockRepository = mocks.MockContactRepository{}
	)

	mockRepository.On("PutItem", mockContact).Return(nil)

	p := processor.New(&mockRepository)

	contact, err := p.Process(mockRequest)

	// Assertions
	require.NoError(t, err, "an error was returned")
	assert.Equal(t, mockContact, contact)
}

func TestProcessor_Process_FailInternalError(t *testing.T) {
	var (
		mockRepository = mocks.MockContactRepository{}
	)

	mockRepository.On("PutItem", mockContact).Return(&ddb.InternalError{
		Op:  "PutItem",
		Err: errors.New("something bad happened"),
	})

	p := processor.New(&mockRepository)

	_, err := p.Process(mockRequest)

	// Assertions
	require.Error(t, err, "should return an error")
	assert.Equal(t, true, errors.As(err, &mockInternalError))
}
