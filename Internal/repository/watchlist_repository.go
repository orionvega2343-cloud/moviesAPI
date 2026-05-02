package repository

import (
	"MobeReviewAPI/Internal/models"

	"github.com/jmoiron/sqlx"
)

type WatchlistRepo struct {
	repo *sqlx.DB
}

func NewWatchlistRepo(db *sqlx.DB) *WatchlistRepo {
	return &WatchlistRepo{repo: db}
}

func (r *WatchlistRepo) CreateWatchlist(w *models.Watchlist) error {
	_, err := r.repo.Exec(`INSERT INTO watchlist (user_id,movie_id) VALUES ($1,$2)`, w.UserID, w.MovieID)
	if err != nil {
		return err
	}
	return nil
}

func (r *WatchlistRepo) GetWatchlistById(id int) (*models.Watchlist, error) {
	var w models.Watchlist
	err := r.repo.QueryRow(`SELECT user_id, movie_id FROM watchlist WHERE id = $1`, id).Scan(&w.UserID, &w.MovieID)
	if err != nil {
		return nil, err
	}
	return &w, nil
}
