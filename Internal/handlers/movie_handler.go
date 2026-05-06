package handlers

import (
	"MobeReviewAPI/Internal/models"
	"MobeReviewAPI/Internal/service"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
)

type MovieHandler struct {
	DB *sqlx.DB
}

func NewMovieHandler(db *sqlx.DB) *MovieHandler {
	return &MovieHandler{DB: db}
}
func (h *MovieHandler) CreateMovie(w http.ResponseWriter, r *http.Request) {
	var m models.Movie
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = service.CreateMovieService(h.DB, &m)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = json.NewEncoder(w).Encode(m)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func (h *MovieHandler) GetMovieById(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	convId, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	res, err := service.GetByIdService(h.DB, convId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}
