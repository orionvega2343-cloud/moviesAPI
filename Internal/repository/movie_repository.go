package repository

import (
	"MobeReviewAPI/Internal/models"

	"github.com/jmoiron/sqlx"
)

type MovieRepo struct {
	repo *sqlx.DB
}

func NewMovieRepo(db *sqlx.DB) *MovieRepo {
	return &MovieRepo{repo: db}
}

func (r *MovieRepo) CreateMovie(m *models.Movie) error {
	_, err := r.repo.Exec(`INSERT INTO movies (title,year,genre,director) VALUES ($1,$2,$3,$4)`, m.Title, m.Year, m.Genre, m.Director)
	if err != nil {
		return err
	}
	return nil
}

func (r *MovieRepo) GetById(id int) (*models.Movie, error) {
	var m models.Movie
	err := r.repo.QueryRow(`SELECT id,title,year,genre,director FROM movies WHERE id=$1`, id).Scan(&m.Id, &m.Title, &m.Year, &m.Genre, &m.Director)
	if err != nil {
		return nil, err
	}
	return &m, nil
}
