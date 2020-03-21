package handlers

import (
	"github.com/slham/toolbelt/l"
	"net/http"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	l.Debug(r.Context(), "Skole!")
	_, _ = w.Write([]byte("Skole!"))
}
