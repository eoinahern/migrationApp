package routes

import "net/http"

// SessionHandler handler the session control e.g. login to
//youtube maintain access while downloading / uploading content
type SessionHandler struct {
}

func (s *SessionHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("hello from session"))
}
