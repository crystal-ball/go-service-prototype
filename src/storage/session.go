package storage

import (
	"os"

	"github.com/gorilla/sessions"
)

// SessionStore provides cookie session storage
var SessionStore = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))
