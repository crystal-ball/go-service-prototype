package routes

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Index handles requests to / with a welcome message
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")
}

// Hello handles requests to /hello with a custom name message
func Hello(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
    name := vars["name"]

	fmt.Fprintf(w, "hello, %s!\n", name)
}