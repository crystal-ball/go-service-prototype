package main

import (
	"io"
	"net/http"

	"github.com/gorilla/mux"

	"go-service-prototype/src/routes"
)

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	io.WriteString(w, `{"alive": true}`)
}

// InitRouter sets up the service routes
func InitRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/health", healthCheckHandler)
	router.HandleFunc("/", routes.Index)
	router.HandleFunc("/hello/{name}", routes.Hello)

	return router
}
