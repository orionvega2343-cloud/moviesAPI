package models

type Movie struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Year     int    `json:"year"`
	Genre    string `json:"genre"`
	Director string `json:"director"`
}
