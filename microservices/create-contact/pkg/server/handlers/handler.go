package handlers

import (
	"net/http"
	"seb7887/create-contact/pkg/contacts"
)

type Handler interface {
	GetHealth(w http.ResponseWriter, r *http.Request)
	CreateContact(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	service contacts.ContactService
}

func New(s contacts.ContactService) Handler {
	return &handler{
		service: s,
	}
}
