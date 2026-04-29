package repository

import (
	"MobeReviewAPI/Internal/models"

	"github.com/jmoiron/sqlx"
)

func CreateReviews(db *sqlx.DB, r models.Review) error {
	_, err := db.Exec(`INSERT INTO reviews (user_id,movie_id,rating,text) VALUES ($1,$2,$3,$4)`, r.UserID, r.MovieID, r.Rating, r.Text)
	if err != nil {
		return err
	}
	return nil
}

func GetReviewsById(db *sqlx.DB, id int) (models.Review, error) {
	var r models.Review
	err := db.QueryRow(`SELECT user_id,movie_id,rating,text FROM reviews WHERE id = $1`, id).Scan(&r.UserID, &r.MovieID, &r.Rating, &r.Text)
	if err != nil {
		return r, err
	}
	return r, nil
}
