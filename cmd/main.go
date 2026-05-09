package main

import (
	"MobeReviewAPI/Internal/config"
	"MobeReviewAPI/Internal/db"
	"MobeReviewAPI/Internal/handlers"
	"MobeReviewAPI/Internal/middlewares"
	"MobeReviewAPI/Internal/repository"
	"MobeReviewAPI/Internal/service"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func main() {
	cfg := config.MustLoad()
	database := db.Connect(&cfg.DB)

	uRepo := repository.NewUserRepo(database)

	uService := service.NewUserService(uRepo, cfg.JWT.Secret)

	uHandler := handlers.NewUserHandler(database, uService)
	mHandler := handlers.NewMovieHandler(database)
	rHandler := handlers.NewReviewsHandler(database)
	wHandler := handlers.NewWatchlistHandler(database)

	r := chi.NewRouter()
	r.Use(middlewares.LoggerMiddleware)
	r.Post("/auth/register", uHandler.CreateUser)
	r.Post("/auth/login", uHandler.Login)

	r.Group(func(r chi.Router) {
		r.Use(func(next http.Handler) http.Handler {
			return middlewares.AuthMiddleware(next, cfg.JWT.Secret)
		})
		r.Get("/movie/{id}", mHandler.GetMovieById)
		r.Post("/movie", mHandler.CreateMovie)

		r.Get("/reviews/{id}", rHandler.GetReviewsById)
		r.Post("/reviews", rHandler.CreateReviews)

		r.Get("/watchlist/{id}", wHandler.GetWatchlistById)
		r.Post("/watchlist", wHandler.NewWatchlist)

	})
	err := http.ListenAndServe(":"+strconv.Itoa(cfg.Server.Port), r)
	if err != nil {
		log.Fatal(err)
	}
}
