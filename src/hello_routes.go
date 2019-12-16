package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Index handles requests to / with a welcome message
func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

// Hello handles requests to /hello with a custom name message
func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}
