package repository

import (
	"MobeReviewAPI/Internal/models"

	"github.com/jmoiron/sqlx"
)

func CreateWatchlist(db *sqlx.DB, w *models.Watchlist) error {
	_,err := db.Exec(`INSERT INTO watchlist (user_id,movie_id) VALUES ($1,$2)`, w.UserID, w.MovieID)
	if err != nil {
		return err
	}
	return nil
}

func GetWatchlistById(db *sqlx.DB, id int) (*models.Watchlist, error) {\
	var w models.Watchlist
	err := db.QueryRow(`SELECT user_id, movie_id FROM watchlist WHERE id = $1`, id).Scan(&w.UserID, &w.MovieID)
	if err != nil {
		return nil, err
	}
	return &w, nil
}