package server

import (
	"context"
	"github.com/gorilla/mux"
	"net/http"
	"seb7887/create-contact/pkg/contacts"
	"seb7887/create-contact/pkg/server/handlers"
)

const (
	HealthEndpoint = "/health"
	ContactsEndpoint = "/contacts"
)

type Server interface {
	Router() http.Handler
	Serve() error
	Shutdown(ctx context.Context) error
}

func New(addr string) http.Server {
	return http.Server{
		Addr: addr,
		Handler: router(),
	}
}

func router() http.Handler {
	var (
		r = mux.NewRouter()
		repository = contacts.NewContactsRepository()
		service = contacts.NewContactsService(repository)
		handler = handlers.New(service)
	)

	r.HandleFunc(HealthEndpoint, handler.GetHealth).Methods(http.MethodGet)
	r.HandleFunc(ContactsEndpoint, handler.CreateContact).Methods(http.MethodPost)

	return r
}

