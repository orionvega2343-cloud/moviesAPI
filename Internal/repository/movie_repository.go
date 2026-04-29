package repository

import (
	"MobeReviewAPI/Internal/models"

	"github.com/jmoiron/sqlx"
)

func CreateMovie(db *sqlx.DB, m *models.Movie) error {
	_, err := db.Exec(`INSERT INTO movies (title,year,genre,director) VALUES ($1,$2,$3,$4)`, m.Title, m.Year, m.Genre, m.Director)
	if err != nil {
		return err
	}
	return nil
}

func GetById(db *sqlx.DB, id int) (*models.Movie, error) {
	var m models.Movie
	err := db.QueryRow(`SELECT id,title,year,genre,director FROM movies WHERE id=$1`, id).Scan(&m.Id, &m.Title, &m.Year, &m.Genre, &m.Director)
	if err != nil {
		return nil, err
	}
	return &m, nil
}
