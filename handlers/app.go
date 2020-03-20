package handlers

import (
	"github.com/slham/toolbelt"
	"net/http"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	toolbelt.Debug(r.Context(), "Skole!")
	_, _ = w.Write([]byte("Skole!"))
}
