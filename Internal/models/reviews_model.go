package models

type Review struct {
	Id      int    `json:"id"`
	UserID  int    `json:"user_id"`
	MovieID int    `json:"movie_id"`
	Rating  int    `json:"rating"`
	Text    string `json:"text"`
}
