package service

import (
	"github.com/gorilla/mux"

	"go-service-prototype/src/middleware"
	"go-service-prototype/src/routes"
)

// InitRouter sets up the service routes
func InitRouter() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/healthcheck", routes.Healthcheck)

	router.HandleFunc("/account/create", routes.CreateAccount)
	router.HandleFunc("/account", routes.GetAccount)

	router.HandleFunc("/auth/login", routes.LoginAccount)

	// --- Middleware

	router.Use(middleware.Logger)

	return router
}
