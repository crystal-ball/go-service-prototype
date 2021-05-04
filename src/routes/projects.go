package routes

import (
	"encoding/json"
	"log"
	"net/http"
)

type Dummy struct {
	Success bool
}

// ListProjects returns all projects for the user
func ListProjects(w http.ResponseWriter, r *http.Request) {
	log.Println("Get account")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(struct{ Success bool }{Success: true})
}

// CreateProject creates a new project
func CreateProject(w http.ResponseWriter, r *http.Request) {
	log.Println("Create project")

	res := Dummy{
		Success: true,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(res)

}

// GetProject returns the details for the requested project
func GetProject(w http.ResponseWriter, r *http.Request) {
	log.Println("Get account")

	res := Dummy{
		Success: true,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(res)
}

// UpdateProject updates the project
func UpdateProject(w http.ResponseWriter, r *http.Request) {
	log.Println("Get account")

	res := Dummy{
		Success: true,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(res)
}

// DeleteProject deletes the project
func DeleteProject(w http.ResponseWriter, r *http.Request) {
	log.Println("Get account")

	res := Dummy{
		Success: true,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(res)
}
