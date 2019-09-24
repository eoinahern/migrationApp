package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/eoinahern/migration_app/routes"
)

func main() {
	port := os.Getenv("PORT")
	http.Handle("/session", &routes.SessionHandler{})
	http.Handle("/upload", &routes.UploadHandler{})
	http.Handle("/download", &routes.DownloadHandler{})
	log.Printf("Listening on port %s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}
