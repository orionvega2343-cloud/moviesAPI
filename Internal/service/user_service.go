package service

import (
	"MobeReviewAPI/Internal/models"
	"MobeReviewAPI/Internal/repository"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	repo   *repository.UserRepo
	Secret string
}

func NewUserService(repo *repository.UserRepo, secret string) *UserService {
	return &UserService{repo: repo, Secret: secret}
}

type User struct {
	UserID int    `db:"user_id"`
	Email  string `db:"email"`
	jwt.RegisteredClaims
}

func (u *UserService) Register(w *models.Register) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(w.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	w.Password = string(hash)
	err = u.repo.CreateUser(w)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserService) Login(l *models.Login, password string) (string, error) {
	res, err := u.repo.GetByEmail(*l)
	if err != nil {
		return "", err
	}
	err = bcrypt.CompareHashAndPassword([]byte(res.Password), []byte(password)) //Сравниваем пароли
	if err != nil {
		return "", err
	}
	var c User
	c.UserID = res.Id
	c.Email = l.Email
	c.ExpiresAt = jwt.NewNumericDate(time.Now().Add(24 * time.Hour))
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	tokenString, err := token.SignedString([]byte(u.Secret))
	if err != nil {
		return "", err
	}
	return tokenString, nil

}
