package mocks

import (
	ddb "github.com/seb7887/aws-lib/dynamodb"
	"github.com/stretchr/testify/mock"
)

type MockContactRepository struct {
	mock.Mock
}

func (m *MockContactRepository) GetByPk(id string) (ddb.Contact, error) {
	args := m.Called(id)
	if args.Get(1) != nil {
		return args.Get(0).(ddb.Contact), args.Error(1)
	}
	return args.Get(0).(ddb.Contact), nil
}

func (m *MockContactRepository) PutItem(contact ddb.Contact) error {
	args := m.Called(contact)
	if args.Get(0) != nil {
		return args.Get(0).(error)
	}
	return nil
}
