package db

import (
	"MobeReviewAPI/Internal/config"
	"fmt"

	"github.com/jmoiron/sqlx"
)

// Connect DB
func Connect(cfg *config.DBConfig) *sqlx.DB {
	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable", cfg.Host, cfg.User, cfg.Pass, cfg.Name, cfg.Port)
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		panic(err)
	}
	return db
}
