package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bradgignac/fortune-api/db"

	"goji.io"
	"goji.io/pat"
)

// Handler is an http.Handler for serving the API
type Handler struct {
	*goji.Mux
	db *db.Database
}

// NewHandler creates an api.Handler that serves from the provided database.
func NewHandler(db *db.Database) *Handler {
	api := Handler{Mux: goji.NewMux(), db: db}
	api.HandleFunc(pat.Get("/fortunes"), api.list)
	api.HandleFunc(pat.Get("/fortunes/:id"), api.get)
	api.HandleFunc(pat.Get("/random"), api.random)
	return &api
}

func (h *Handler) list(w http.ResponseWriter, r *http.Request) {
	fortunes := h.db.List()

	w.Header().Add("Content-Type", "application/json")
	enc := json.NewEncoder(w)
	enc.Encode(fortunes)
}

func (h *Handler) get(w http.ResponseWriter, r *http.Request) {
	id := pat.Param(r, "id")
	f, err := h.db.Get(id)
	if err == db.ErrMissingFortune {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	enc := json.NewEncoder(w)
	enc.Encode(f)
}

func (h *Handler) random(w http.ResponseWriter, r *http.Request) {
	id, err := h.db.Random()
	if err == db.ErrEmptyDatabase {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	http.Redirect(w, r, fmt.Sprintf("/fortunes/%s", id), 302)
}
