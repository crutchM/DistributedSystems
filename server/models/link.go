package models

type Link struct {
	Id     int    `json:"id"`
	Url    string `json:"link" db:"link"`
	Status string `json:"status"`
}
