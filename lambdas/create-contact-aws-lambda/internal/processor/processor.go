package processor

import (
	"uala/go-workshop/internal/repository"
	"uala/go-workshop/pkg/dto"
)

type Processor interface {
	Setup()
	Process(req dto.Request) (dto.Contact, error)
}

type LambdaProcessor struct {
	ContactRepository repository.Repository
}

func NewProcessor(r repository.Repository) Processor {
	return &LambdaProcessor{
		ContactRepository: r,
	}
}

func (p *LambdaProcessor) Setup() {
	// Singleton
	if p.ContactRepository == nil {
		p.ContactRepository = &repository.LambdaRepository{}
	}
	p.ContactRepository.Setup()
}

func (p *LambdaProcessor) Process(req dto.Request) (dto.Contact, error) {
	contact := dto.Contact{
		ID:        "",
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Status:    "CREATED",
	}

	item, err := p.ContactRepository.Insert(contact)
	if err != nil {
		return dto.Contact{}, err
	}

	return item, nil
}
