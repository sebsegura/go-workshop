package processor

import (
	ddb "github.com/seb7887/aws-lib/dynamodb"
	"uala/go-workshop/pkg/dto"
)

type Processor interface {
	Process(req dto.Request) (ddb.Contact, error)
}

type LambdaProcessor struct {
	ContactRepository ddb.ContactsRepository
}

const (
	Created = "CREATED"
)

func New(r ddb.ContactsRepository) Processor {
	return &LambdaProcessor{
		ContactRepository: r,
	}
}

func (p *LambdaProcessor) Process(req dto.Request) (ddb.Contact, error) {
	contact := ddb.Contact{
		ID:        "",
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Status:    Created,
	}

	err := p.ContactRepository.PutItem(contact)
	if err != nil {
		return ddb.Contact{}, err
	}

	return contact, nil
}
