package handlers

import (
	"log"
	"net/http"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	log.Println("Skole!")
	_, _ = w.Write([]byte("Skole!"))
}
