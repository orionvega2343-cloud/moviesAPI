package models

type Watchlist struct {
	Id      int `json:"id"`
	UserID  int `json:"user_id"`
	MovieID int `json:"movie_id"`
}
