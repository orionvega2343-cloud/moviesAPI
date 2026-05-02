package service

import (
	"MobeReviewAPI/Internal/models"
	"MobeReviewAPI/Internal/repository"

	"github.com/jmoiron/sqlx"
)

func CreateWatchlist(db *sqlx.DB,w *models.Watchlist) error {
	repo := repository.NewWatchlistRepo(db)
	err := repo.CreateWatchlist(w)
	if err != nil {
		return err
	}
	return nil
}

func GetWatchlist(db *sqlx.DB,id int) (*models.Watchlist, error) {
	repo := repository.NewWatchlistRepo(db)
	w, err := repo.GetWatchlistById(id)
	if err != nil {
		return nil, err
	}
	return w, nil
}