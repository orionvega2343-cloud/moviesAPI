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

type ReviewsHandler struct {
	DB *sqlx.DB
}

func NewReviewsHandler(db *sqlx.DB) *ReviewsHandler {
	return &ReviewsHandler{DB: db}
}

func (h *ReviewsHandler) CreateReviews(w http.ResponseWriter, r *http.Request) {
	var rev models.Review
	err := json.NewDecoder(r.Body).Decode(&rev)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = service.CreateReviewsService(h.DB, &rev)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(rev)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *ReviewsHandler) GetReviewsById(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	convId, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	res, err := service.GetReviewsById(h.DB, convId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
