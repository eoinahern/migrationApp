package main

import (
	"net/http"

	"github.com/eoinahern/migration_app/routes"
)

func main() {
	http.Handle("/session", &routes.SessionHandler{})
	http.ListenAndServe(":8080", nil)
}
