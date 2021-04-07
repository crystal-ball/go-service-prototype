package routes

import (
	"io"
	"net/http"
)

// HealthCheck returns an alive message when the service is up
func Healthcheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	io.WriteString(w, `{"alive": true}`)
}
