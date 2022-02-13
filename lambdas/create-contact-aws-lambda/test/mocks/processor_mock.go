package mocks

import (
	ddb "github.com/seb7887/aws-lib/dynamodb"
	"github.com/stretchr/testify/mock"
	"uala/go-workshop/pkg/dto"
)

type MockProcessor struct {
	mock.Mock
}

func (m *MockProcessor) Process(req dto.Request) (ddb.Contact, error) {
	args := m.Called(req)
	if args.Get(1) != nil {
		return args.Get(0).(ddb.Contact), args.Error(1)
	}
	return args.Get(0).(ddb.Contact), nil
}
