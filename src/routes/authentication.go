package routes

import (
	"io"
	"net/http"

	"go-service-prototype/src/storage"
)

func LoginAccount(w http.ResponseWriter, r *http.Request) {
	session, _ := storage.SessionStore.Get(r, "session")

	session.Values["account_id"] = "12345"

	err := session.Save(r, w)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	io.WriteString(w, `{"success": true}`)
}
