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

type WatchlistHandler struct {
	DB *sqlx.DB
}

func NewWatchlistHandler(db *sqlx.DB) *WatchlistHandler {
	return &WatchlistHandler{DB: db}
}

func (h *WatchlistHandler) NewWatchlist(w http.ResponseWriter, r *http.Request) {
	var wl models.Watchlist
	err := json.NewDecoder(r.Body).Decode(&wl)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = service.CreateWatchlist(h.DB, &wl)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = json.NewEncoder(w).Encode(wl)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *WatchlistHandler) GetWatchlistById(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	res, err := service.GetWatchlist(h.DB, id)
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
