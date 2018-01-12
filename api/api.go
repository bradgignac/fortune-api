package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/bradgignac/fortune-api/fortune"

	"goji.io"
	"goji.io/pat"
)

type Handler struct {
	*goji.Mux
	db *fortune.Database
}

func NewHandler(db *fortune.Database) *Handler {
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
	fortune := h.db.Get(id)

	w.Header().Add("Content-Type", "application/json")
	enc := json.NewEncoder(w)
	enc.Encode(fortune)
}

func (h *Handler) random(w http.ResponseWriter, r *http.Request) {
	id := h.db.Random()
	http.Redirect(w, r, fmt.Sprintf("/fortunes/%s", id), 302)
}
