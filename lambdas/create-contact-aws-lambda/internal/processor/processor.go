package processor

import (
	db "github.com/seb7887/aws-lib/dynamodb"
	"uala/go-workshop/internal/utils"
	"uala/go-workshop/pkg/dto"
)

type Processor interface {
	Process(req dto.Request) (db.Contact, error)
}

type LambdaProcessor struct {
	ContactRepository db.ContactsRepository
}

const (
	Created = "CREATED"
)

func New(r db.ContactsRepository) Processor {
	return &LambdaProcessor{
		ContactRepository: r,
	}
}

func (p *LambdaProcessor) Process(req dto.Request) (db.Contact, error) {
	contact := db.Contact{
		ID:        utils.GenerateUUID(),
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Status:    Created,
	}

	err := p.ContactRepository.PutItem(contact)
	if err != nil {
		return db.Contact{}, err
	}

	return contact, nil
}
