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

	router.HandleFunc("/auth/signin", routes.LoginAccount).Methods("POST")
	router.HandleFunc("/auth/signout", routes.LoginAccount).Methods("POST")

	router.HandleFunc("/projects", routes.ListProjects).Methods("GET")
	router.HandleFunc("/projects/{projectID}", routes.GetProject).Methods("GET")
	router.HandleFunc("/projects/{projectID}", routes.CreateProject).Methods("POST")
	router.HandleFunc("/projects/{projectID}", routes.UpdateProject).Methods("PUT")
	router.HandleFunc("/projects/{projectID}", routes.DeleteProject).Methods("DELETE")

	// --- Middleware

	router.Use(middleware.Logger)

	return router
}
