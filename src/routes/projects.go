package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Dummy struct {
	Success bool
}

// ListProjects returns all projects for the user
func ListProjects(w http.ResponseWriter, r *http.Request) {
	log.Println("Get projects")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(struct{ Success bool }{Success: true})
}

// CreateProject creates a new project
func CreateProject(w http.ResponseWriter, r *http.Request) {
	log.Println("Create project")

	// body := r.Body.Read()

	res := Dummy{
		Success: true,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(res)

}

type ProjectsRes struct {
	Success bool   `json:"success"`
	ID      string `json:"id"`
}

// GetProject returns the details for the requested project
func GetProject(w http.ResponseWriter, r *http.Request) {
	log.Println("Get single project")

	vars := mux.Vars(r)

	res := ProjectsRes{
		Success: true,
		ID:      vars["projectID"],
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
