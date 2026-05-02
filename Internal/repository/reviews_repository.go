package repository

import (
	"MobeReviewAPI/Internal/models"

	"github.com/jmoiron/sqlx"
)

type ReviewsRepo struct {
	repo *sqlx.DB
}

func NewReviewsRepo(db *sqlx.DB) *ReviewsRepo {
	return &ReviewsRepo{repo: db}
}

func (r *ReviewsRepo) CreateReviews(a *models.Review) error {
	_, err := r.repo.Exec(`INSERT INTO reviews (user_id,movie_id,rating,text) VALUES ($1,$2,$3,$4)`, a.UserID, a.MovieID, a.Rating, a.Text)
	if err != nil {
		return err
	}
	return nil
}

func (r *ReviewsRepo) GetReviewsById(id int) (*models.Review, error) {
	var a models.Review
	err := r.repo.QueryRow(`SELECT user_id,movie_id,rating,text FROM reviews WHERE id = $1`, id).Scan(&a.UserID, &a.MovieID, &a.Rating, &a.Text)
	if err != nil {
		return nil, err
	}
	return &a, nil
}
