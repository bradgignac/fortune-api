package api

import "net/http"

type Handler struct {
}

func (a *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("fortune-api"))
}
