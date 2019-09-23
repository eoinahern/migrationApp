package routes

import "net/http"

// SessionHandler handler the session control e.g. login to
//youtube maintain access while downloading / uploading content
type SessionHandler struct {
}

//DownloadHandler : download an temp store files
type DownloadHandler struct {
}

//UploadHandler : upload to other platforms
type UploadHandler struct {
}

func (s *SessionHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("hello from session"))
}

func (d *DownloadHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("hello from session"))
}

func (u *UploadHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("hello from session"))
}
