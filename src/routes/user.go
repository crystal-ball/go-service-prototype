package routes

import (
	"encoding/json"
	"log"
	"net/http"
)

type Account struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// CreateAccount creates a new account with provided the email and password
func CreateAccount(w http.ResponseWriter, r *http.Request) {
	log.Println("Create account")

	var account Account
	json.NewDecoder(r.Body).Decode(&account)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(account)
}

// GetAccount returns the account for the provided email and password
func GetAccount(w http.ResponseWriter, r *http.Request) {
	log.Println("Get account")

	account := Account{
		Name:  "Hedge",
		Email: "hedge@mail.com",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(account)
}
