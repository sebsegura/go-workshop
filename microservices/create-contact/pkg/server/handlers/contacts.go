package handlers

import (
	"encoding/json"
	"net/http"
	"seb7887/create-contact/pkg/models"
)

func (h *handler) CreateContact(w http.ResponseWriter, r *http.Request) {
	var req models.CreateContactRequest
	_ = json.NewDecoder(r.Body).Decode(&req)

	if !validateRequest(req) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write([]byte("Invalid request"))
		return
	}

	contact, err := h.service.CreateContact(req)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		b, _ := json.Marshal(err)
		_, _ = w.Write(b)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	b, _ := json.Marshal(contact)
	_, _ = w.Write(b)
}

func validateRequest(req models.CreateContactRequest) bool {
	return req.FirstName != "" && req.LastName != ""
}
