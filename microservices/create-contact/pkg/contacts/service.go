package contacts

import (
	"seb7887/create-contact/pkg/models"
	"seb7887/create-contact/pkg/uuid"
)

const (
	Created = "CREATED"
)

type ContactService interface {
	CreateContact(req models.CreateContactRequest) (models.Contact, error)
}

type contactService struct {
	r ContactRepository
}

func NewContactsService(r ContactRepository) ContactService {
	return &contactService{
		r,
	}
}

func (s *contactService) CreateContact(req models.CreateContactRequest) (models.Contact, error) {
	contact := models.Contact{
		Uuid:      uuid.GetUUID(),
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Status:    Created,
	}

	err := s.r.PutItem(contact)

	return contact, err
}
