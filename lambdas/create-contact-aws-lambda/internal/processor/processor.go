package processor

import (
	"seb7887/create-contact/internal/repository"
	"seb7887/create-contact/pkg/dto"
)

type Processor interface {
	Setup()
	Process(req dto.Request) (dto.Contact, error)
}

type LambdaProcessor struct {
	Repository repository.Repository
}

func (p *LambdaProcessor) Setup() {
	if p.Repository == nil {
		p.Repository = &repository.ContactRepository{}
	}
	p.Repository.Setup()
}

func (p *LambdaProcessor) Process(req dto.Request) (dto.Contact, error) {
	contact := dto.Contact{
		ID:        "",
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Status:    "CREATED",
	}

	_, err := p.Repository.Insert(contact)
	if err != nil {
		return dto.Contact{}, err
	}

	return contact, nil
}
