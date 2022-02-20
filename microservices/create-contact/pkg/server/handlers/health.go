package handlers

import (
	"fmt"
	"net/http"
)

func (h *handler) GetHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	_, _ = fmt.Fprintf(w, "ok")
}
