package repository

import (
	"MobeReviewAPI/Internal/models"

	"github.com/jmoiron/sqlx"
)

type UserRepo struct {
	repo *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) *UserRepo {
	return &UserRepo{repo: db}
}

func (r *UserRepo) CreateUser(u *models.Register) error {

	_, err := r.repo.Exec(`INSERT INTO users (email, name, password) VALUES ($1, $2, $3)`, u.Email, u.Name, u.Password)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepo) GetByEmail(l models.Login) (*models.Register, error) {
	user := &models.Register{}
	err := r.repo.QueryRow(`SELECT id,email,password FROM users WHERE email=$1`, l.Email).Scan(&user.Id, &user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return user, err
}
