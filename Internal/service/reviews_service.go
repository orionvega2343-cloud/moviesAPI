package service

import (
	"MobeReviewAPI/Internal/models"
	"MobeReviewAPI/Internal/repository"

	"github.com/jmoiron/sqlx"
)

func CreateReviewsService(db *sqlx.DB, r *models.Review) error {
	repo := repository.NewReviewsRepo(db)
	err := repo.CreateReviews(r)
	if err != nil {
		return err
	}
	return nil
}

func GetReviewsById(db *sqlx.DB, id int) (*models.Review, error) {
	repo := repository.NewReviewsRepo(db)
	review, err := repo.GetReviewsById(id)
	if err != nil {
		return nil, err
	}
	return review, nil
}
