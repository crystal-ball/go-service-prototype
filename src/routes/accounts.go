package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Account struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// CreateAccount creates a new account with provided the email and password
func CreateAccount(w http.ResponseWriter, r *http.Request) {
	var account Account
	json.NewDecoder(r.Body).Decode(&account)

	fmt.Fprintf(w, "Account for %s created for %s email", account.Name, account.Email)
}

// GetAccount returns the account for the provided email and password
func GetAccount(w http.ResponseWriter, r *http.Request) {
	account := Account{
		Name:  "Hedge",
		Email: "hedge@mail.com",
	}

	json.NewEncoder(w).Encode(account)
}
