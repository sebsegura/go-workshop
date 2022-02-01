package processor

import "uala/go-workshop/internal/repository"

// TODO: interface & struct
type Processor interface {
	Setup()
	Process() error
}

type LambdaProcessor struct {
	ContactRepository repository.Repository
}

// TODO: Setup processor -> instanciar repository
func (p *LambdaProcessor) Setup() {
	// Singleton
	if p.ContactRepository == nil {
		p.ContactRepository = &repository.LambdaRepository{}
	}
	p.ContactRepository.Setup()
}

// TODO: Process
func (p *LambdaProcessor) Process() error {
	return nil
}