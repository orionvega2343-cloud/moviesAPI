package service

import (
	"MobeReviewAPI/Internal/models"
	"MobeReviewAPI/Internal/repository"

	"github.com/jmoiron/sqlx"
)

func CreateMovieService(db *sqlx.DB, m *models.Movie) error {
	repo := repository.NewMovieRepo(db)
	err := repo.CreateMovie(m)
	if err != nil {
		return err
	}
	return nil
}

func GetByIdService(db *sqlx.DB, id int) (*models.Movie, error) {
	repo := repository.NewMovieRepo(db)
	movie, err := repo.GetById(id)
	if err != nil {
		return nil, err
	}
	return movie, nil
}
