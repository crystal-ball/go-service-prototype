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

	router.HandleFunc("/user/account", routes.GetAccount).Methods("GET")
	router.HandleFunc("/user/account", routes.CreateAccount).Methods("POST")

	router.HandleFunc("/auth/login", routes.LoginAccount).Methods("POST")
	router.HandleFunc("/auth/logout", routes.LogoutAccount).Methods("POST")

	router.HandleFunc("/projects", routes.ListProjects).Methods("GET")
	router.HandleFunc("/projects", routes.CreateProject).Methods("POST")
	router.HandleFunc("/projects/{projectID}", routes.GetProject).Methods("GET")
	router.HandleFunc("/projects/{projectID}", routes.UpdateProject).Methods("PATCH")
	router.HandleFunc("/projects/{projectID}", routes.DeleteProject).Methods("DELETE")

	// --- Middleware

	router.Use(middleware.Logger)

	return router
}
