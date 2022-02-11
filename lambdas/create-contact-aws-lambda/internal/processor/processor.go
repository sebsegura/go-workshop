package processor

import (
	"uala/go-workshop/internal/repository"
	"uala/go-workshop/internal/utils"
	"uala/go-workshop/pkg/dto"
)

type Processor interface {
	Process(req dto.Request) (dto.Contact, error)
}

type LambdaProcessor struct {
	ContactRepository repository.Repository
}

const (
	Created = "CREATED"
)

func New(r repository.Repository) Processor {
	return &LambdaProcessor{
		ContactRepository: r,
	}
}

func (p *LambdaProcessor) Process(req dto.Request) (dto.Contact, error) {
	contact := dto.Contact{
		ID:        utils.GenerateUUID(),
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Status:    Created,
	}

	item, err := p.ContactRepository.Insert(contact)
	if err != nil {
		return dto.Contact{}, err
	}

	return item, nil
}
