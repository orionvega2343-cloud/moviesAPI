package repository

import (
	"MobeReviewAPI/Internal/models"

	"github.com/jmoiron/sqlx"
)

func CreateUser(db *sqlx.DB, u *models.Register) error {

	_, err := db.Exec(`INSERT INTO users (email, name, password) VALUES ($1, $2, $3)`, u.Email, u.Name, u.Password)
	if err != nil {
		return err
	}
	return nil
}

func GetByEmail(db *sqlx.DB, l models.Login) (*models.Register, error) {
	user := &models.Register{}
	err := db.QueryRow(`SELECT email,password FROM users WHERE email=$1`, l.Email).Scan(&user.Email, &user.Password)
	if err != nil {
		return nil, err
	}
	return user, err
}
